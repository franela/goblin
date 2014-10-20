// +build windows

package goblin

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return false
}
