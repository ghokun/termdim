// +build android darwin dragonfly freebsd ios linux netbsd openbsd

package termdim

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func getSize(fd int) (width, height int, err error) {
	size, err := unix.IoctlGetWinsize(fd, syscall.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	return int(size.Col), int(size.Row), nil
}
