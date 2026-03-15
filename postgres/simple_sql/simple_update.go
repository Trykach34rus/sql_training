package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateBook(ctx context.Context,conn *pgx.Conn, books Book) error {
	sqlQuery := `
	UPDATE books 
	SET title=$1,author=$2,description=$3,release=$4,read=$5,addBook=$6,readFullBook=$7
	WHERE id=$8 ;
	`

	_,err := conn.Exec(ctx,sqlQuery,books.Title,books.Author,books.Description,books.Release,books.Read,books.AddBook,books.ReadFullBook,books.ID)

	return  err
}