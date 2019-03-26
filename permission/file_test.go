package permission

import (
	"log"
	"os"
	"syscall"
	"testing"
)

func TestWriteX(t *testing.T) {
	err := WriteX("/")
	if err == nil {
		log.Fatal("Expect error")
	}
	if err.(*os.PathError).Err != syscall.EACCES {
		log.Fatalln("Expect syscall.EACCES, got:", err.(*os.PathError).Err)
	}
}
