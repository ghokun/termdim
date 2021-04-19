// +build android darwin dragonfly freebsd ios linux netbsd openbsd

package termdim

import (
	"os/exec"
	"testing"

	"github.com/creack/pty"
)

func TestGetSizeUnix(t *testing.T) {
	cmd := exec.Command("ls")
	ptyx, err := pty.StartWithSize(cmd, &pty.Winsize{
		Rows: 50,
		Cols: 100,
		X:    0,
		Y:    0,
	})
	defer func() { _ = ptyx.Close() }()
	
	if err != nil {
		t.Fatal(err)
	}
	width, height, err := GetSize(int(ptyx.Fd()))
	if err != nil {
		t.Fatal(err)
	}
	if width != 100 {
		t.Fatalf("%d != %d", width, 100)
	}
	t.Logf("Terminal width: %d", width)
	if height != 50 {
		t.Fatalf("%d != %d", height, 50)
	}
	t.Logf("Terminal height: %d", height)
}
