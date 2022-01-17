package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
)

type Transaction interface {
	WithLogger(l logging.Logger) Transaction

	Query(query string, args ...interface{}) (rows *Rows, err error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Commit() error
	MustRollBack(entailedError string)
	RollbackIfNotDone()
	Prepare(query string) (*Stmt, error)
	IsInitialized() bool
	SetInternalTransaction(tx *sql.Tx)
	SetLogger(l logging.Logger)
}

type sqlTransaction struct {
	l          logging.Logger
	tx         *sql.Tx
	OnCommit   func() error
	OnRollback func() error
}

func (t sqlTransaction) WithLogger(l logging.Logger) Transaction {
	t.l = l
	return &t
}

//nolint:sqlclosecheck
func (t *sqlTransaction) Query(query string, args ...interface{}) (rows *Rows, err error) {
	defer processError(&err)

	tsBeforeRequest := time.Now().UTC()
	sqlRows, err := t.tx.Query(query, args...)
	if err != nil || sqlRows.Err() != nil {
		return
	}

	latency := math.Floor(time.Now().UTC().Sub(tsBeforeRequest).Seconds()*1000) / 1000
	t.l.WithField("query", query).WithField("fields", args).WithField("db_latency", latency).Debug("db profiling")
	rows = &Rows{
		rows: sqlRows,
	}

	return
}

func (t *sqlTransaction) Commit() (err error) {
	defer processError(&err)

	if t == nil {
		return nil
	}

	if t.tx == nil {
		return ErrorProducerUnspecified.New("Nil tx commit")
	}

	tsBeforeRequest := time.Now().UTC()
	err = t.tx.Commit()

	latency := math.Floor(time.Now().UTC().Sub(tsBeforeRequest).Seconds()*1000) / 1000
	t.l.WithField("query", "commit").WithField("db_latency", latency).Debug("db profiling")

	return
}

func (t *sqlTransaction) MustRollBack(entailed string) {
	if t == nil || t.tx == nil {
		return
	}

	if err := t.tx.Rollback(); err != nil {
		errMsg := fmt.Sprintf("entailed rollback error: %s, rollback err: %s", entailed, err.Error())
		err = ErrorProducerUnspecified.New(errMsg)
		panic(err.Error())
	}
}

func (t *sqlTransaction) RollbackIfNotDone() {
	if t == nil || t.tx == nil {
		return
	}

	err := t.tx.Rollback()
	if err != nil {
		if errors.Is(err, sql.ErrTxDone) {
			return
		}

		panic(err.Error())
	}
}

func (t *sqlTransaction) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	defer processError(&err)

	tsBeforeRequest := time.Now().UTC()

	result, err = t.tx.Exec(query, args...)

	latency := math.Floor(time.Now().UTC().Sub(tsBeforeRequest).Seconds()*1000) / 1000
	t.l.WithField("query", query).WithField("fields", args).WithField("db_latency", latency).Debug("db profiling")

	return
}

//nolint:sqlclosecheck
func (t *sqlTransaction) Prepare(query string) (result *Stmt, err error) {
	defer processError(&err)

	stmt, err := t.tx.Prepare(query)
	if err != nil {
		return
	}

	result = &Stmt{
		l:     t.l,
		stmt:  stmt,
		query: query,
	}

	return
}

func (t *sqlTransaction) IsInitialized() bool {
	return t.tx != nil
}

func (t *sqlTransaction) SetInternalTransaction(tx *sql.Tx) {
	t.tx = tx
}

func (t *sqlTransaction) SetLogger(l logging.Logger) {
	t.l = l
}
