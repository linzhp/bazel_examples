package socket

import (
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestSocket(t *testing.T) {
	f, err := ioutil.TempFile(os.Getenv("TEST_TMPDIR"), "unix-socket")
	if err != nil {
		t.Error(err)
	}
	os.Remove(f.Name())

	ln2, err := net.Listen("unix", f.Name())
	if err != nil {
		t.Error(err)
	}
	defer ln2.Close()
}