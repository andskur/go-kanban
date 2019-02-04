package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/andskur/kanban-board/config"
	"github.com/andskur/kanban-board/internal/server"
)

// Application represent application structure
type Application struct {
	*config.Config
}

// NewApplication create new application instance
func NewApplication(cfg *config.Config) (*Application, error) {
	app := &Application{
		Config: cfg,
	}

	return app, nil
}

// StartServer start http web server
func (app *Application) StartServer() {
	// register handler
	http.HandleFunc("/", server.BoardHandler(app.Config))

	addr := strings.Join([]string{app.Host, app.Port}, ":")
	// Listen connections
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
