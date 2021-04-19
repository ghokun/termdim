// +build windows

package termdim

import (
	"testing"

	"github.com/mattn/go-tty"
)

func TestGetSizeWindows(t *testing.T) {
	ttyx, err := tty.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer ttyx.Close()

	width, height, err := GetSize(int(ttyx.Output().Fd()))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Terminal width: %d", width)
	t.Logf("Terminal height: %d", height)
}
