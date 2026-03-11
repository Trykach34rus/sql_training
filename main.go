package main

import (
	"context"
	"fmt"
	"sql_training/postgres/simple_connection"
	"sql_training/postgres/simple_sql"
	"time"
)

type Book struct{
	title string
	author string
	description string
	release int
	read bool
	addBook time.Time
	readFullBook *time.Time
	
}



func main() {
	ctx := context.Background()

	book := Book{
    title: "Война и мир",
    author: "Толстой",
    description: "что-то там про войну и мир",
    release: 1869,
    read: false,
    addBook: time.Now(),
}

	conn,err := simple_connection.CreateConnection(ctx)
	if err != nil{
		panic(err)
	}
	
	if err := simple_sql.CreateTable(ctx,conn);err != nil{
		panic(err)
	}

 if err := simple_sql.InseartBook(ctx, conn, book.title, book.author, book.description, book.release, book.read, book.addBook); err !=nil{
	panic(err)
 }

	fmt.Println("Работает хреновина")



}