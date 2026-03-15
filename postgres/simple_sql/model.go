package simple_sql

import "time"

type Book struct {
	ID           int
	Title        string
	Author       string
	Description  string
	Release      int
	Read         bool
	AddBook      time.Time
	ReadFullBook *time.Time
}