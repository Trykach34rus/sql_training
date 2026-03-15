package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ListPages(ctx context.Context,conn *pgx.Conn, N int)([]Book,error) {

	sqlQuery := `
	SELECT id,title,author,description,release,read,addBook,readFullBook FROM books
	ORDER BY ID ASC
	LIMIT $1
	OFFSET $2
	`
	var allBooks []Book
	page := 0

	for {
		offset := page *N
		rows,err :=conn.Query(ctx,sqlQuery,N,offset)
		if err != nil {
			return nil,err
		}

		var pagesBook []Book

		for rows.Next(){
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
			if err !=nil{
				return nil,err
			}

			pagesBook = append(pagesBook, b)

		}
		if rows.Err() !=nil{
			return nil,rows.Err()
		}
		rows.Close()

		if len(pagesBook) == 0 {
			break
		}

		fmt.Println("Страница",page+1)

		for _,b := range pagesBook{
			fmt.Println(b.ID,b.Title,b.Author)
		}

		if len(pagesBook) < N {
    fmt.Println("Это последняя страница")
}

		allBooks = append(allBooks, pagesBook...)
		page++
	}

	return allBooks,nil


}