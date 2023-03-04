package global

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const (
	contextTxKey = "ContextTxKey"
)

type DBController struct {
	db *sql.DB
}

func NewDBController(strConnectDB string) DBController {
	dbPoll, err := sql.Open("pgx", strConnectDB)
	if err != nil {
		log.Fatal(err)
	}

	err = dbPoll.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return DBController{db: dbPoll}
}

func (pc DBController) CloseConnect() {
	if err := pc.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (pc DBController) NewTxContext(ctx context.Context) (context.Context, error) {
	tx, err := pc.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, contextTxKey, tx), nil
}

func (pc DBController) CommitTx(ctx context.Context) error {
	tx, ok := ctx.Value(contextTxKey).(*sql.Tx)
	if !ok {
		return fmt.Errorf("don't run commit - no open transactions")
	}

	return tx.Commit()
}

func (pc DBController) RollbackTx(ctx context.Context) {
	tx, ok := ctx.Value(contextTxKey).(*sql.Tx)
	if !ok {
		log.Println("RollbackTx err: no open transaction")
		return
	}

	if err := tx.Rollback(); err != nil {
		log.Printf("RollbackTx err: %s\n", err)
	}
}

func (pc DBController) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, ok := ctx.Value(contextTxKey).(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("don't run exec - no open transactions")
	}

	return tx.ExecContext(ctx, query, args...)
}

func (pc DBController) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tx, ok := ctx.Value(contextTxKey).(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("don't run Query - no open transactions")
	}

	return tx.QueryContext(ctx, query, args...)
}

func (pc DBController) QueryRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	tx, ok := ctx.Value(contextTxKey).(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("don't run QueryRow - no open transactions")
	}

	return tx.QueryRowContext(ctx, query, args...), nil
}
