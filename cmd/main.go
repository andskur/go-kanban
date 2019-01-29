package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/andskur/kanban-board/internal/app"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(app.Board)
}
