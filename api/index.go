package handler

import (
	"net/http"
	"sync"

	"tixora/internal/config"
	"tixora/internal/server"
)

var (
	app  http.Handler
	once sync.Once
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		cfg := config.LoadEnv()
		db := config.ConnectDatabase(cfg)
		app = server.New(db, cfg)
	})

	app.ServeHTTP(w, r)
}
