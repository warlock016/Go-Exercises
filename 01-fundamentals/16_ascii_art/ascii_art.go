package asciiart

import "strings"

// DrawBox creates a rectangular box using the specified character.
// The box will have 'width' characters horizontally and 'height' rows vertically.
// If width or height is less than 2, returns an empty string.
//
// Example: DrawBox(5, 3, '*') produces:
//
//	*****
//	*   *
//	*****
func DrawBox(width, height int, char rune) string {
	// TODO(human): Implement box drawing
	if height < 2 || width < 2 {
		return ""
	}

	x := string(char)
	var box strings.Builder

	for h := range height {
		if h == 0 || h == height-1 {
			box.WriteString(strings.Repeat(x, width))
			box.WriteString("\n")
		} else {
			for w := range width {
				switch w {
				case 0:
					box.WriteString(x)
				case width - 1:
					box.WriteString(x)
					box.WriteString("\n")
				default:
					box.WriteString(" ")
				}
			}
		}
	}

	return box.String()
}

// DrawDiamond creates a diamond pattern with n rows in the top half.
// The diamond uses '*' characters and is centered with spaces.
// If n is less than 1, returns an empty string.
//
// Example: DrawDiamond(3) produces:
//
//	  *
//	 ***
//	*****
//	 ***
//	  *
func DrawDiamond(n int) string {
	// TODO(human): Implement diamond drawing

	if n < 1 {
		return ""
	}

	var diamond strings.Builder

	maxRows := 2*n - 1
	currSymbolCount := 1
	currSpaceCount := n - 1
	ascend := true

	for r := 0; r < maxRows; r++ { // number of rows calculated via (2x - 1)
		switch ascend {
		case true:
			diamond.WriteString(strings.Repeat(" ", currSpaceCount))
			diamond.WriteString(strings.Repeat("*", currSymbolCount))
			diamond.WriteString("\n")

			if currSpaceCount == 0 {
				ascend = false
				// currSpaceCount += 1
				// currSymbolCount -= 2
				continue
			}
			currSpaceCount -= 1
			currSymbolCount += 2
		case false:
			currSpaceCount += 1
			currSymbolCount -= 2
			diamond.WriteString(strings.Repeat(" ", currSpaceCount))
			diamond.WriteString(strings.Repeat("*", currSymbolCount))
			diamond.WriteString("\n")
		}
	}
	return diamond.String()
}

// DrawChessboard creates an n×n chessboard pattern using filled (█) and empty (░) blocks.
// The pattern alternates based on position: if (row+col) is even, use █, otherwise use ░.
// If n is less than 1, returns an empty string.
//
// Example: DrawChessboard(4) produces:
//
//	█░█░
//	░█░█
//	█░█░
//	░█░█
func DrawChessboard(n int) string {
	// TODO(human): Implement chessboard drawing
	if n < 1 {
		return ""
	}

	white := '█'
	black := '░'

	var board strings.Builder

	for i := range n {
		for j := range n {
			switch (i+j)%2 == 0 {
			case true:
				board.WriteString(string(white))
			case false:
				board.WriteString(string(black))
			}

			if j+1 == n {
				board.WriteString("\n")
			}
		}
	}

	return board.String()
}
