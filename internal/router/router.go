package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"multifinance-go/internal/config"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	cfg := config.NewViper()
	Router(r)
	port := cfg.GetString("web.port")
	appName := cfg.GetString("app.name")
	log := config.NewLogger(cfg)

	log.WithFields(logrus.Fields{
		"appName": appName,
		"port":    port,
	}).Info("Starting the application")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.WithField("error", err).Fatal("Failed to start the server")
	}
}

func Router(r *mux.Router) {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	_ = config.NewDatabase(viperConfig, log)
}
