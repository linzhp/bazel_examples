package gomock

import (
	"github.com/mattn/go-sqlite3"
	"fmt"
)


type sshClient interface {
	Close() error
}

func Print() {
	fmt.Println(sqlite3.Version())
}
