// +build windows

package termdim

import (
	"golang.org/x/sys/windows"
)

func getSize(fd int) (width, height int, err error) {
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(fd), &info); err != nil {
		return -1, -1, err
	}
	return info.Window.Right - info.Window.Left + 1, info.Window.Bottom - info.Window.Top + 1, nil
}
