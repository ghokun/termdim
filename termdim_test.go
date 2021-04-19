package termdim

import (
	"os/exec"
	"testing"

	"github.com/creack/pty"
)

const (
	expectedWidth  = 100
	expectedHeight = 50
)

func TestGetSize(t *testing.T) {
	cmd := exec.Command("dir")
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
