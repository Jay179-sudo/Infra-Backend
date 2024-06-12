package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const MaxAllotedTime int = 10

type config struct {
	port int    // :4000
	env  string // (Development|Staging|Production)
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	environment := os.Getenv("ENVIRONMENT")
	portAddress := os.Getenv("PORT")

	// setting the default values
	cfg.env = "Development"
	cfg.port = 4000

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	if environment == "Staging" || environment == "Production" {
		cfg.env = environment
	}
	port, _ := strconv.Atoi(portAddress)
	// default to port 4000 in case of input error
	if port > 1023 {
		cfg.port = port
	}

	// client, err := connectDB(cfg.db.connString)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatalf("Could not disconnect from DB")
	// 	}
	// }()
	var app application

	app.config = cfg
	app.logger = logger
	// app.mongo = client

	// setup MongoDB Index to Email Field
	// err = app.createIndex()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
