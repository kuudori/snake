package gameboard

const (
	DefaultX1 = 1
	DefaultY1 = 1
	DefaultX2 = 65
	DefaultY2 = 10
)

type Board struct {
	X1, Y1, X2, Y2 int
}

func NewBoard() Board {
	return Board{
		X1: DefaultX1, Y1: DefaultX1,
		X2: DefaultX2, Y2: DefaultY2,
	}
}

func (b *Board) Width() int {
	return b.X2 - b.X1
}

func (b *Board) Height() int {
	return b.Y2 - b.Y1
}

func (b *Board) Update(x1, y1, x2, y2 int) {
	b.X1, b.Y1, b.X2, b.Y2 = x1, y1, x2, y2
}

func (b *Board) Reset() {
	b.Update(DefaultX1, DefaultY1, DefaultX2, DefaultY2)
}

func IsBoardOk(x1, y1, x2, y2 int) bool {
	return x2-x1 > 0 && y2-y1 > 0
}
