package database

import (
	"database/sql"
	"regexp"
	"strconv"
	"sync"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"

	"github.com/jmoiron/sqlx"
)

var (
	mu    sync.Mutex
	sqlDB *sql.DB
)

func GetDatabase(config Config, l logging.Logger) (result *sqlDatabase, err error) {
	defer processError(&err)

	mu.Lock()
	defer mu.Unlock()

	if sqlDB == nil {
		_db, err := sql.Open("postgres", config.ConnectionDSN)
		if err != nil {
			return nil, err
		}
		sqlDB = _db
	}

	result = &sqlDatabase{
		client:  sqlDB,
		xClient: sqlx.NewDb(sqlDB, "postgres"),
		logger:  l,
	}

	return
}

func NewSqlTransaction(db Database) (result *sql.Tx, err error) {
	defer processError(&err)

	result, err = db.Begin()
	if err != nil {
		return nil, err
	}

	return
}

func InitIfNot(db Database, withTx Transaction, l logging.Logger) (err error) {
	defer processError(&err)

	if withTx == nil || !withTx.IsInitialized() {
		var sqlTx *sql.Tx
		sqlTx, err = NewSqlTransaction(db)
		if err != nil {
			return
		}
		withTx.SetInternalTransaction(sqlTx)
		withTx.SetLogger(l)
	} else {
		withTx.SetLogger(l)
	}

	return
}

func SliceToInterfaceSlice(source []string) (result []interface{}) {
	for _, elem := range source {
		result = append(result, elem)
	}
	return
}

func BuildInQuery(argNumber int) (result string) {
	if argNumber == 0 {
		return ""
	}

	for i := 0; i < argNumber; i++ {
		result += "$" + strconv.Itoa(i+1) + ","
	}
	return result[:len(result)-1]
}

func CloseStmt(stmt *Stmt, err *error) {
	if stmt != nil {
		if tempErr := stmt.Close(); tempErr != nil && *err == nil {
			*err = tempErr
		}
	}

	if *err != nil {
		return
	}
}

func CloseConnections(stmt *Stmt, rows *Rows, err *error) {
	if stmt != nil {
		if tempErr := stmt.Close(); tempErr != nil && *err == nil {
			*err = tempErr
		}
	}

	if rows != nil {
		if tempErr := rows.Close(); tempErr != nil {
			*err = tempErr
		}
	}

	if *err != nil {
		return
	}
}

func EmptyTransaction(l logging.Logger) Transaction {
	return &sqlTransaction{
		l: l,
	}
}

// Rebind a query from the default bindtype (QUESTION) to the postgress bindtype(DOLLAR).
// e.g. Rebind("where some = ? and thing = ?") == "where some = $1 and thing = $2".
func Rebind(query string) string {
	return sqlx.Rebind(sqlx.DOLLAR, query)
}

var likeEscaping = regexp.MustCompile(`([_%])`)

// LikeSubstr creates pattern for the LIKE/ILIKE expression to find corresponding substring.
// It also escape an underscore (_) and a percent sign(%).
func LikeSubstr(str string) string {
	if str == "" {
		return "%"
	}

	escaped := likeEscaping.ReplaceAllString(str, `\$1`)

	return "%" + escaped + "%"
}

// In expands slice values in args, returning the modified query string
// and a new arg list that can be executed by a database. The `query` should
// use the `?` bindVar.  The return value uses the `?` bindVar.
func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}
