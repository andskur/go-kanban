package server

import (
	"html/template"
	"net/http"

	"github.com/andskur/kanban-board/config"
	"github.com/andskur/kanban-board/internal/services/kanban_board"
)

// BoardHandler handling main request to server and response as a html Kanban Board
func BoardHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch board data
		board, err := kanban.NewBoard(cfg)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		// Init frontend
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		if err := tmpl.Execute(w, board); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
