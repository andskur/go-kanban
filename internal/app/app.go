package application

import (
	"github.com/andskur/kanban-board/config"
	"github.com/andskur/kanban-board/internal/kanban_board"
)

// Application represent application structure
type Application struct {
	*config.Config
	Board kanban.Board
}

// NewApplication create new application instance
func NewApplication() (*Application, error) {
	app := &Application{
		Config: config.InitConfig(),
	}
	err := app.initBoard()
	if err != nil {
		return nil, err
	}

	return app, nil
}

// initBoard init Kanban board for current application
func (app *Application) initBoard() error {
	board, err := kanban.NewBoard(app.Config)
	if err != nil {
		return err
	}

	app.Board = *board
	return nil
}
