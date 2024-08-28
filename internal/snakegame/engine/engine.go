package engine

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"math/rand"
	"snake/config"
	"snake/internal/snakegame/gameboard"
	"snake/internal/snakegame/snake"
	"snake/internal/ui"
	"time"
)

type Game struct {
	screen      tcell.Screen
	snake       *snake.Snake
	food        snake.Point
	board       gameboard.Board
	gameRunning bool
	score       int
	inputChan   chan tcell.Event
	ticker      *time.Ticker
	speed       time.Duration
	eventChan   chan tcell.Event
}

func NewGame(s tcell.Screen, board gameboard.Board) *Game {
	g := &Game{
		screen:    s,
		board:     board,
		snake:     snake.NewSnake(board.Width()/2, board.Height()/2),
		inputChan: make(chan tcell.Event),
		speed:     config.DEFAULT_SPEED,
		eventChan: make(chan tcell.Event),
	}
	g.ticker = time.NewTicker(g.speed)
	g.placeFood()
	return g
}

func (g *Game) placeFood() {
	for {
		x := rand.Intn(g.board.Width())
		y := rand.Intn(g.board.Height())
		if !g.snake.Contains(snake.Point{X: x, Y: y}) {
			g.food = snake.Point{X: x, Y: y}
			return
		}
	}
}

func (g *Game) Run() {
	g.gameRunning = true

	eventChan := make(chan tcell.Event)

	go func() {
		for g.gameRunning {
			eventChan <- g.screen.PollEvent()
		}
		close(g.eventChan)
	}()

	g.runGameLoop(eventChan)
}

func (g *Game) runGameLoop(eventChan <-chan tcell.Event) {
	defer g.Stop()

	for g.gameRunning {
		select {
		case ev := <-eventChan:
			if ev != nil {
				g.handleInput(ev)
			}
		case <-g.ticker.C:
			g.Update()
			g.Draw()
		}
	}
}

func (g *Game) Stop() {
	g.ticker.Stop()
	close(g.inputChan)
}

func (g *Game) Update() {
	newHead := g.snake.Body[0]
	newHead.X = (newHead.X + g.snake.Direction.X + g.board.Width()) % g.board.Width()
	newHead.Y = (newHead.Y + g.snake.Direction.Y + g.board.Height()) % g.board.Height()

	for _, bodyPart := range g.snake.Body[1:] {
		if newHead == bodyPart {
			g.gameRunning = false
			return
		}
	}

	if newHead == g.food {
		g.snake.Body = append([]snake.Point{newHead}, g.snake.Body...)
		g.placeFood()
		g.score++
	} else {
		g.snake.Body = append([]snake.Point{newHead}, g.snake.Body[:len(g.snake.Body)-1]...)
	}
}

// Draw Draws snake movement
func (g *Game) Draw() {
	ui.DrawBox(g.screen, g.board.X1, g.board.Y1, g.board.X2, g.board.Y2, ui.GetBlackBoxStyle())
	for i, p := range g.snake.Body {
		ch := '◉'
		if i > 0 {
			ch = '○'
		}
		g.screen.SetContent(g.board.X1+p.X, g.board.Y1+p.Y, ch, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	}

	g.screen.SetContent(g.board.X1+g.food.X, g.board.Y1+g.food.Y, '●', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))

	scoreStr := fmt.Sprintf("Score: %d", g.score)
	for i, ch := range scoreStr {
		startX := g.board.X1 + (g.board.Width()-len(scoreStr))/2
		startY := g.board.Y2 - 5
		g.screen.SetContent(startX+i, startY, ch, nil, ui.GetBlackBoxStyle())
	}
	g.screen.Show()
}

func (g *Game) handleInput(ev tcell.Event) {

	switch ev := ev.(type) {

	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			g.gameRunning = false
			return
		case tcell.KeyUp:
			if g.snake.Direction.Y == 0 {
				g.snake.Direction = snake.Point{X: 0, Y: -1}
			}

		case tcell.KeyDown:
			if g.snake.Direction.Y == 0 {
				g.snake.Direction = snake.Point{X: 0, Y: 1}
			}

		case tcell.KeyLeft:
			if g.snake.Direction.X == 0 {
				g.snake.Direction = snake.Point{X: -1, Y: 0}
			}

		case tcell.KeyRight:
			if g.snake.Direction.X == 0 {
				g.snake.Direction = snake.Point{X: 1, Y: 0}
			}
		case tcell.KeyPgDn:
			g.speed += 10 * time.Millisecond
			if g.speed > 500*time.Millisecond {
				g.speed = 500 * time.Millisecond
			}
			g.ticker.Reset(g.speed)
		case tcell.KeyPgUp:
			g.speed -= 10 * time.Millisecond
			if g.speed < 50*time.Millisecond {
				g.speed = 50 * time.Millisecond
			}
			g.ticker.Reset(g.speed)

		}

	case *tcell.EventResize:
		g.screen.Sync()
	}
}
