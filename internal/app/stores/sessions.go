package stores

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
)

type sessions struct {
	db database.Database
	l  logging.Logger
}

func NewSessionsStore(db database.Database, l logging.Logger) repositories.Sessions {
	return &sessions{
		db: db,
		l:  l,
	}
}

func (s sessions) WithLogger(l logging.Logger) repositories.Sessions {
	s.l = l
	s.db = s.db.WithLogger(l)
	return &s
}

func (s sessions) ByUUID(UUID string) (session models.Session, err error) {
	defer processError(&err)

	stmt, err := s.db.Prepare(`
SELECT uuid,
       user_uuid,
       expiring_at,
       created_at
FROM sessions
WHERE uuid = $1;
	`)
	if err != nil {
		database.CloseStmt(stmt, &err)
		return
	}
	defer database.CloseStmt(stmt, &err)

	row := stmt.QueryRow(UUID)
	err = row.Scan(
		&session.UUID,
		&session.UserUUID,
		&session.ExpiringAt,
		&session.CreatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (s sessions) ByUserUUID(UUID string) (session models.Session, err error) {
	defer processError(&err)

	stmt, err := s.db.Prepare(`
SELECT uuid,
       user_uuid,
       expiring_at,
       created_at
FROM sessions
WHERE user_uuid = $1;
	`)
	if err != nil {
		database.CloseStmt(stmt, &err)
		return
	}
	defer database.CloseStmt(stmt, &err)

	row := stmt.QueryRow(UUID)
	err = row.Scan(
		&session.UUID,
		&session.UserUUID,
		&session.ExpiringAt,
		&session.CreatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (s sessions) Set(session models.Session, withTx database.Transaction) (err error) {
	defer processError(&err)

	err = database.InitIfNot(s.db, withTx, s.l)
	if err != nil {
		return
	}

	query := `
INSERT INTO sessions(uuid, user_uuid, expiring_at, created_at)
VALUES ($1, $2, $3, $4)
ON CONFLICT(user_uuid) DO UPDATE SET uuid=$1,
                                     expiring_at=$3,
                                     created_at=$4;`
	result, err := withTx.Exec(query, session.UUID, session.UserUUID, session.ExpiringAt.UTC(),
		session.CreatedAt.UTC())
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

func (s sessions) Delete(userUUID string, withTx database.Transaction) (err error) {
	defer processError(&err)

	err = database.InitIfNot(s.db, withTx, s.l)
	if err != nil {
		return
	}

	query := `DELETE FROM sessions WHERE user_uuid=$1`
	result, err := withTx.Exec(query, userUUID)
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
