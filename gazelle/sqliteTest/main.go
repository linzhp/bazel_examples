package main

import (
	"fmt"

	"github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(sqlite3.Version())
}
