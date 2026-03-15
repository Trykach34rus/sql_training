package main

import (
	"context"
	"fmt"
	"sql_training/postgres/simple_connection"
	"sql_training/postgres/simple_sql"
	"time"
)





func main() {
	ctx := context.Background()

	newBook := simple_sql.Book{
    Title: "Мёртвые души",
    Author: "Гогаль",
    Description: "что-то там про Мёртвые души",
    Release: 1869,
    Read: false,
    AddBook: time.Now(),
}

	conn,err := simple_connection.CreateConnection(ctx)
	if err != nil{
		panic(err)
	}
	
	if err := simple_sql.CreateTable(ctx,conn);err != nil{
		panic(err)
	}

	books,err := simple_sql.GetBooks(ctx,conn)
	if err !=nil{
	  panic(err)
	}

  for _, b := range books {
		if b.ID == 2 {
			b.Title = "Сказка о царе султане"
			b.Author = "Пушкин"
			b.Description = "Что там о царе султане"
			b.Release = 1800
			b.Read = false
			b.AddBook = time.Now()
			now := time.Now()
			b.ReadFullBook = &now

			if err := simple_sql.UpdateBook(ctx,conn,b);err != nil{
			panic(err)
		  }
		break
	  }
  }

	allBooks,err := simple_sql.ListPages(ctx,conn,5)
	if err != nil{
		panic(err)
	}

	fmt.Println("Все книги полученные через пагинацию:",len(allBooks))

 if err := simple_sql.InseartBook(ctx, conn,newBook); err !=nil{
	panic(err)
 }
 
 if err := simple_sql.DeleteBook(ctx,conn,2); err != nil{
	panic(err)
 }



	fmt.Println("Работает хреновина")



}