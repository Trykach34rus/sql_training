package main

import (
	"context"
	"fmt"
	"sql_training/postgres/simple_connection"
)

func main() {
	ctx := context.Background()

	conn,err := simple_connection.CreateConnection(ctx)
	if err != nil{
		panic(err)
	}
	conn.Ping(ctx)
	fmt.Println("Работает хреновина")

}