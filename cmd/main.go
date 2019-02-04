package main

import (
	"github.com/andskur/kanban-board/config"
	"github.com/andskur/kanban-board/internal/app"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

	go func() {
		log.Printf("listen and serve on %s\n", addr)
		http.ListenAndServe(addr, nil)
	}()

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-quit
	log.Println("Shutdown Server ...")
}
