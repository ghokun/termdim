package termdim

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSize(t *testing.T) {
	width, height, err := GetSize(int(os.Stdout.Fd()))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(width)
	fmt.Println(height)
}
