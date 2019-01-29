package main

import (
	"fmt"
	"github.com/andskur/kanban-board/internal/app"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/andskur/kanban-board/config"
)

func main() {
	config := config.InitConfig()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app, err := application.NewApplication(config)
		if err != nil {
			log.Fatal(err)
		}
		tmpl, _ := template.ParseFiles("views/index.html")
		tmpl.Execute(w, app.Board)
	})
	addr := strings.Join([]string{config.Host, config.Port}, ":")

	fmt.Println("Server is listening...")
	http.ListenAndServe(addr, nil)
}
