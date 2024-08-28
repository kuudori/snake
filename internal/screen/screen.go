package screen

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"snake/config"
	"snake/internal/snakegame/gameboard"
	"snake/internal/ui"
)

// InitializeStartScreen Initializes the screen
func InitializeStartScreen() (tcell.Screen, chan bool) {
	done := make(chan bool)
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()

	return s, done
}

// ResetStartScreen Resets the screen to the initial state (menu)
func ResetStartScreen(s tcell.Screen, board gameboard.Board, done chan bool) {
	s.Clear()
	ui.DrawBox(s, board.X1, board.Y1, board.X2, board.Y2, ui.GetBlackBoxStyle())
	ui.DrawAnimatedText(s, board, config.LogoTitle, ui.TopCenter)
	go ui.BlinkText(s, config.StartTitle, board, ui.BottomCenter, done)
}

// ShowHelpScreen Shows help info
func ShowHelpScreen(s tcell.Screen, board gameboard.Board) {
	ui.DrawBox(s, board.X1, board.Y1, board.X2, board.Y2, ui.GetBlackBoxStyle())
	ui.DrawAnimatedText(s, board, config.HelpTitle, ui.TopCenter)
}
