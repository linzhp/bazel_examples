package gomock

import (
	"github.com/mattn/go-sqlite3"
	_ "github.com/bazelbuild/bazel-gazelle/cmd/gazelle"
	"fmt"
)


type sshClient interface {
	Close() error
}

func Print() {
	fmt.Println(sqlite3.Version())
}
