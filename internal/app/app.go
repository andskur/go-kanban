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
	err := app.FetchBoard()
	if err != nil {
		return nil, err
	}

	return app, nil
}

// FetchBoard fetch data for Kanban board for current application
func (app *Application) FetchBoard() error {
	board, err := kanban.NewBoard(app.Config)
	if err != nil {
		return err
	}

	app.Board = *board
	return nil
}
