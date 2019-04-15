package workingdir

import (
	"log"
	"os"
	"testing"
)

var outsideWd, _ = os.Getwd()

func TestWorkingDir(t *testing.T) {
	insideWd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if outsideWd != insideWd {
		log.Fatalf("Expect no working directory change, got\n\t%s and\n\t%s", outsideWd, insideWd)
	}
}
