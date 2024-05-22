package router

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"multifinance-go/internal/config"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	cfg := config.NewViper()

	port := cfg.GetString("web.port")
	appName := cfg.GetString("app.name")

	// Log application start
	log.WithFields(log.Fields{
		"appName": appName,
		"port":    port,
	}).Info("Application is starting")

	// Start the server and log any errors
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.WithFields(log.Fields{
			"port": port,
			"err":  err,
		}).Fatal("Failed to start server")
	}
}
