package main

import (
	"log"
	"tixora/pkg/config"
	"tixora/pkg/server"
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