package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/andskur/kanban-board/internal/app"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	board := app.Board

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("views/index.html")
		tmpl.Execute(w, board)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
