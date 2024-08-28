package ui

import (
	"github.com/gdamore/tcell/v2"
	"snake/internal/snakegame/gameboard"
	"strings"
	"time"
)

type Position int

const (
	TopLeft Position = iota
	TopCenter
	TopRight
	CenterLeft
	Center
	CenterRight
	BottomLeft
	BottomCenter
	BottomRight
)

// getPositionCoords Returns coordinates for position
func getPositionCoords(pos Position, board gameboard.Board, textWidth, textHeight int) (int, int) {
	x1, y1, x2, y2 := board.X1, board.Y1, board.X2, board.Y2
	width, height := x2-x1+1, y2-y1+1

	var x, y int

	switch pos {
	case TopLeft:
		x, y = x1, 2
	case TopCenter:
		x, y = x1+(width-textWidth)/2, y1+1
	case TopRight:
		x, y = x2-textWidth+1, y1+1
	case CenterLeft:
		x, y = x1, y1+(height-textHeight)/2
	case Center:
		x, y = x1+(width-textWidth)/2, y1+(height-textHeight)/2
	case CenterRight:
		x, y = x2-textWidth+1, y1+(height-textHeight)/2
	case BottomLeft:
		x, y = x1+2, y2-textHeight-1
	case BottomCenter:
		x, y = x1+(width-textWidth)/2, y2-textHeight-1
	case BottomRight:
		x, y = x2-textWidth-2, y2-textHeight-1
	}

	return x, y
}

// DrawBox Draws a box
func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

}

// GetBlackBoxStyle Returns default box style with black background and while foreground
func GetBlackBoxStyle() tcell.Style {
	return tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
}

// DrawText Draws a text
func DrawText(s tcell.Screen, x, y int, ch rune) {
	s.SetContent(x, y, ch, nil, GetBlackBoxStyle())
}

// DrawAnimatedText Makes a drawing animation for text
func DrawAnimatedText(s tcell.Screen, board gameboard.Board, text string, pos Position) {
	lines := strings.Split(strings.TrimSpace(text), "\n")

	titleWidth := 0
	for _, line := range lines {
		if len(line) > titleWidth {
			titleWidth = len(line)
		}
	}

	startX, startY := getPositionCoords(pos, board, titleWidth, len(lines))

	for y, line := range lines {
		for x, ch := range line {
			DrawText(s, startX+x, startY+y, ch)
			s.Show()
			time.Sleep(2 * time.Millisecond)
		}
		time.Sleep(2 * time.Millisecond)
	}

	time.Sleep(300 * time.Millisecond)
}

// BlinkText Makes a blinking animation for text
func BlinkText(s tcell.Screen, text string, board gameboard.Board, pos Position, done chan bool) {
	textWidth := len(text)
	startX, startY := getPositionCoords(pos, board, textWidth, 1)

	for {
		select {
		case <-done:
			for i := range text {
				s.SetContent(startX+i, startY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorReset))
			}
			return
		default:
			for i, ch := range text {
				DrawText(s, startX+i, startY, ch)
			}
			s.Show()
			time.Sleep(500 * time.Millisecond)

			for i := range text {
				s.SetContent(startX+i, startY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorBlack))
			}
			s.Show()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
