package stores

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
	"time"
)

type users struct {
	db database.Database
	l  logging.Logger
}

func NewUsersStore(db database.Database, l logging.Logger) repositories.Users {
	return &users{
		db: db,
		l:  l,
	}
}

func (u users) WithLogger(l logging.Logger) repositories.Users {
	u.l = l
	u.db = u.db.WithLogger(l)
	return &u
}

func (u users) ByUUID(UUID string) (user models.User, err error) {
	defer processError(&err)

	stmt, err := u.db.Prepare(`
SELECT uuid,
       version,
       first_name,
       last_name,
       bio,
       age,
       gender,
       interests,
       birth_date,
       email
FROM users
WHERE uuid = $1;
	`)
	if err != nil {
		database.CloseStmt(stmt, &err)
		return
	}
	defer database.CloseStmt(stmt, &err)

	row := stmt.QueryRow(UUID)
	err = row.Scan(
		&user.UUID,
		&user.Version,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.Age,
		&user.Gender,
		&user.Interests,
		&user.BirthDate,
		&user.Email,
	)
	if err != nil {
		return
	}

	return
}

func (u users) Set(user models.User, withTx database.Transaction) (err error) {
	defer processError(&err)

	err = database.InitIfNot(u.db, withTx, u.l)
	if err != nil {
		return
	}

	var birthDate *time.Time
	if user.BirthDate != nil {
		temp := user.BirthDate.UTC()
		birthDate = &temp
	}

	query := `
INSERT INTO users(uuid, version, first_name, last_name, bio, age, gender, interests, birth_date, email, password_hash)
VALUES ($1, 1, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT(uuid) DO UPDATE SET version=users.version + 1,
                                updated_at=$3,
                                first_name=$4,
                                last_name=$5,
                                bio=$6,
                                age=$7,
                                gender=$8,
                                interests=$9,
                                birth_date=$10,
                                email=$11,
                                password_hash=$12;`
	result, err := withTx.Exec(query,
		user.UUID,
		user.Version,
		time.Now().UTC(),
		user.FirstName,
		user.LastName,
		user.Bio,
		user.Age,
		user.Gender,
		user.Interests,
		birthDate,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		return
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return
	}

	if ra != 1 {
		err = ErrorProducerInconsistent.New("wrong version")
		return
	}

	return
}

func (u users) SetDeleted(user models.User, deletedAt time.Time, withTx database.Transaction) (err error) {
	defer processError(&err)

	err = database.InitIfNot(u.db, withTx, u.l)
	if err != nil {
		return
	}

	query := `UPDATE users SET version=users.version+1, deleted_at=$2 WHERE uuid=$1 AND users.version=$3;`

	result, err := withTx.Exec(query, user.UUID, deletedAt.UTC(), user.Version)
	if err != nil {
		return
	}

	var ra int64
	ra, err = result.RowsAffected()
	if err != nil {
		return
	}

	if ra != 1 {
		err = ErrorProducerInconsistent.New("wrong version")
		return
	}

	return
}
