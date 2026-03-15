package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)



func GetBooks(ctx context.Context, conn *pgx.Conn) ([]Book,error) {
	sqlQuery := `
	SELECT id,title,author,description,release,read,addBook,readFullBook FROM books
	ORDER BY ID ASC
	LIMIT 10 
	OFFSET 0
	`

	rows,err := conn.Query(ctx,sqlQuery)

	if err != nil{
		return nil,err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var b Book
		err := rows.Scan(
			&b.ID,
      &b.Title,
      &b.Author,
      &b.Description,
      &b.Release,
      &b.Read,
      &b.AddBook,
      &b.ReadFullBook,

        )
		if err !=nil {
			return nil,err
		}		
		books = append(books,b)
	}

	if rows.Err() !=nil {
		return nil,rows.Err()
	}


	return books,nil
}