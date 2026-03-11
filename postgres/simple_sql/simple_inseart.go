package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)



func InseartBook(ctx context.Context,
	conn *pgx.Conn,
	title,author,description string,
	release int,
	read bool,
	addBook time.Time) error {

		

		sqlQuery := `
		INSERT INTO books (title,author,description,release,read,addBook)
		VALUES ($1,$2,$3,$4,$5,$6)
		`

		_,err := conn.Exec(ctx,sqlQuery,title,author,description,release,read,addBook)

		return err

}