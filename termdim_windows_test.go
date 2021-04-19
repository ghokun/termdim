// +build windows

package termdim

import (
	"testing"

	"github.com/mattn/go-tty"
)

func TestGetSizeWindows(t *testing.T) {
	tty, err := tty.Open()
	if err != nil {
		t.Fatal(err)
	}
	winsize := <-tty.SIGWINCH()
	expectedWidth := winsize.W
	expectedHeight := winsize.H

	width, height, err := GetSize(int(tty.Output().Fd()))
	if err != nil {
		t.Fatal(err)
	}
	if width != expectedWidth {
		t.Fatalf("%d != %d", width, expectedWidth)
	}
	t.Logf("Terminal width: %d", width)
	if height != expectedHeight {
		t.Fatalf("%d != %d", height, expectedHeight)
	}
	t.Logf("Terminal height: %d", height)
}
