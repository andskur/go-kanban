package main

import (
	"log"

	"github.com/andskur/kanban-board/config"
	applicaiton "github.com/andskur/kanban-board/internal"
)

func main() {
	conf := config.InitConfig()
	app, err := applicaiton.NewApplication(conf)
	if err != nil {
		log.Fatal(err)
	}
	app.StartServer()
}
