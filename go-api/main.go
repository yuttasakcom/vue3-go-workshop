package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port int
}

type app struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func (app *app) serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Okay    bool   `json:"okay"`
			Message string `json:"message"`
		}

		payload.Okay = true
		payload.Message = "Hello World"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	})

	app.infoLog.Printf("Starting server on port %d", app.config.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), nil)
}

func main() {
	var cfg config
	cfg.port = 8081

	infoLog := log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := app{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
