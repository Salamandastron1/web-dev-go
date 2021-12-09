package session

import (
	"errors"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// (1) need to check that a in-mem is reset
	// (2) check that file system storage is reset
	// got := New()
	// want

}

func TestnewFileStore(t *testing.T) {
	newFileStore(make(inMem))

	_, err := os.Stat("userMap.json")
	if errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}
}
