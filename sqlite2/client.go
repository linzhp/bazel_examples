package sqlite2

import (
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type sshClient interface {
	Close() error
}

func Print() {
	fmt.Println(sqlite3.Version())
}
