// Package t3hterm is currently just a facade for "github.com/olekukonko/ts", which works on Linux, Windows and more.
package t3hterm

import "github.com/olekukonko/ts"

// SizeAndPosition is exactly what it says it is.
type SizeAndPosition struct {
	Width  int
	Height int
	PosX   int
	PosY   int
}

// GetSizeAndPosition talks to the system to get the current size and postion of the terminal.
func GetSizeAndPosition() (ws SizeAndPosition, err error) {
	return getSizeAndPosition()
}

func getSizeAndPosition() (ws SizeAndPosition, err error) {
	size, err := ts.GetSize()
	sp := SizeAndPosition{size.Col(),
		size.Row(),
		size.PosX(),
		size.PosY()}
	return sp, err
}
