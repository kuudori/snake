package main

import (
	"github.com/gdamore/tcell/v2"
	"os"
	"snake/internal/screen"
	"snake/internal/snakegame/engine"
	"snake/internal/snakegame/gameboard"
)

func main() {
	var done chan bool
	var s tcell.Screen

	board := gameboard.NewBoard()
	s, done = screen.InitializeStartScreen()
	screen.ResetStartScreen(s, board, done)

	safeCloseDone := func() {
		if done != nil {
			select {
			case <-done:
			default:
				close(done)
				<-done
			}
			done = nil
		}
	}

	ox, oy := -1, -1
	for {
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				safeCloseDone()
				os.Exit(0)
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				safeCloseDone()
				board.Reset()
				done = make(chan bool)
				screen.ResetStartScreen(s, board, done)
			} else if ev.Key() == tcell.KeyEnter {
				safeCloseDone()
				g := engine.NewGame(s, board)
				g.Run()
				board.Reset()
				done = make(chan bool)
				screen.ResetStartScreen(s, board, done)
			} else if ev.Rune() == 'H' || ev.Rune() == 'h' {
				safeCloseDone()
				screen.ShowHelpScreen(s, board)
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y
				}

			case tcell.ButtonNone:
				if ox >= 1 {
					minX, maxX := min(ox, x), max(ox, x)
					minY, maxY := min(oy, y), max(oy, y)
					ox, oy = -1, -1
					if gameboard.IsBoardOk(minX, minY, maxX, maxY) {
						board.Update(minX, minY, maxX, maxY)
						safeCloseDone()
						s.Clear()
						g := engine.NewGame(s, board)
						g.Run()
						board.Reset()
						done = make(chan bool)
						screen.ResetStartScreen(s, board, done)
					}

				}
			}
		}
	}
}
