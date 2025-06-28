package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Pythonchic/logger"
	"github.com/Pythonchic/tireshop/internal/config"
	"github.com/Pythonchic/tireshop/internal/handlers"
)

var conf config.Config
var err error

func init() {
	conf, err = config.ParseConfig("./config/local.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log := logger.NewLogger()
	log.TimeFormat = "2006/01/02 15:04:05"

	log.Info("Server starting...")

	mainHandler, err := handlers.NewMainHandler("./web/templates/index.html", conf)
	if err != nil {
		log.Error("MainHandler is not initialized", logger.Arg("Error", err))
	} else {
		http.HandleFunc("/", mainHandler)
		log.Info("MainHandler initialized", logger.Arg("Address", "/"))
	}

	infoHandler, err := handlers.NewInfoHandler("./web/templates/info.html", conf)
	if err != nil {
		log.Error("MainHandler is not initialized", logger.Arg("Error", err))
	} else {
		http.HandleFunc("/info", infoHandler)
		log.Info("InfoHandler initialized", logger.Arg("Address", "/info"))
	}

	// Обработчик для статических файлов
	absStaticPath, _ := filepath.Abs("web/static")
	fs := http.FileServer(http.Dir(absStaticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Info("HandlerStaticFiles initialized", logger.Arg("Address", "/static/"))

	absStaticPath, _ = filepath.Abs("storage")
	fs = http.FileServer(http.Dir(absStaticPath))
	http.Handle("/storage/", http.StripPrefix("/storage/", fs))
	log.Info("HandlerStaticFiles initialized", logger.Arg("Address", "/storage/"))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Server started", logger.Arg("Address", conf.Address))

	go func() {
		err := http.ListenAndServe(conf.Address, nil)
		if err != nil {
			log.Error("Failed to start server", logger.Arg("Error", err))
		}
	}()

	<-done
	log.Info("Stopping server...")

	log.Info("Server is stoped")
}
