// +build windows

package termdim

import (
	"testing"

	"github.com/creack/pty"
)

func TestGetSizeWindows(t *testing.T) {
	_, tty, err := pty.Open()
	if err != nil {
		t.Fatal(err)
	}
	width, height, err := GetSize(int(tty.Fd()))
	if err != nil {
		t.Fatal(err)
	}
}
