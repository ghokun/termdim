package termdim

// GetSize returns terminal size
func GetSize(fd int) (width, height int, err error) {
	return getSize(fd)
}
