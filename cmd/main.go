package main

import (
	"log"
	"tixora/internal/config"
	"tixora/internal/server"
)

func main() {
	// load environment variables
	cfg := config.LoadEnv()
	// connect to the database
	db := config.ConnectDatabase(cfg)
	log.Println("✅ Server started successfully")
	// start the server
	server.Start(db, cfg)

}