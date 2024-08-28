package snake

import "github.com/gdamore/tcell/v2"

type Point struct {
	X, Y int
}

type Snake struct {
	Body      []Point
	Direction Point
}

func NewSnake(boardWidth, boardHeight int) *Snake {
	return &Snake{
		Body:      []Point{{X: boardWidth / 2, Y: boardHeight / 2}},
		Direction: Point{X: 1, Y: 0},
	}
}

func (s *Snake) Move(boardWidth, boardHeight int) {
	newHead := Point{
		X: (s.Body[0].X + s.Direction.X + boardWidth) % boardWidth,
		Y: (s.Body[0].Y + s.Direction.Y + boardHeight) % boardHeight,
	}

	s.Body = append([]Point{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Draw(screen tcell.Screen, x int, y int) {
	screen.SetContent(x, y, '■', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
}

func (s *Snake) Grow() {
	newHead := Point{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}

	s.Body = append([]Point{newHead}, s.Body...)
}

func (s *Snake) Contains(p Point) bool {
	for _, bodyPart := range s.Body {
		if bodyPart == p {
			return true
		}
	}
	return false
}
