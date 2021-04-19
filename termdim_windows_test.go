// +build windows

package termdim

import (
	"testing"

	"github.com/mattn/go-tty"
)

func TestGetSizeWindows(t *testing.T) {
	winsize := <-tty.SIGWINCH()
	expectedWidth := winsize.W
	expectedHeight := winsize.H

	ttyx, err := tty.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer ttyx.Close()

	width, height, err := GetSize(int(ttyx.Output().Fd()))
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
