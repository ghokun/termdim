package termdim

import (
	"os/exec"
	"runtime"
	"testing"

	"github.com/creack/pty"
)

const (
	expectedWidth  = 100
	expectedHeight = 50
)

func TestGetSize(t *testing.T) {
	var command string
	if runtime.GOOS == "windows" {
		command = "dir"
	} else {
		command = "ls"
	}
	cmd := exec.Command(command)
	pty, err := pty.StartWithSize(cmd, &pty.Winsize{
		Rows: expectedHeight,
		Cols: expectedWidth,
		X:    0,
		Y:    0,
	})
	if err != nil {
		t.Fatal(err)
	}
	width, height, err := GetSize(int(pty.Fd()))
	if err != nil {
		t.Fatal(err)
	}
	if width != expectedWidth {
		t.Fatalf("%d != %d", width, expectedWidth)
	}
	if height != expectedHeight {
		t.Fatalf("%d != %d", height, expectedHeight)
	}
}
