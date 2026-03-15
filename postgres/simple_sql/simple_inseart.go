package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)



func InseartBook(ctx context.Context,
	conn *pgx.Conn,books Book) error {

		

		sqlQuery := `
		INSERT INTO books (title,author,description,release,read,addBook)
		VALUES ($1,$2,$3,$4,$5,$6)
		`

		_,err := conn.Exec(ctx,sqlQuery,books.Title,books.Author,books.Description,books.Release,books.Read,books.AddBook)

		return err

}