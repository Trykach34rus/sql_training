package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteBook(ctx context.Context,conn *pgx.Conn, ID int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id=$1
	`
	_,err := conn.Exec(ctx,sqlQuery,ID)

	return err
}