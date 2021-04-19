package termdim

import (
	"fmt"
	"testing"

	"github.com/mattn/go-tty"
)

func TestGetSize(t *testing.T) {
	tty, err := tty.Open()
	if err != nil {
		t.Fatal(err)
	}
	tty.Output().Fd()
	width, height, err := GetSize(int(tty.Output().Fd()))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(width)
	fmt.Println(height)
}
