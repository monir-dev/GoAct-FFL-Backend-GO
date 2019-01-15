package main

import (
	"Structure/src/config/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")


	// initialize server
	s := server.NewServer()

	s.Init(port)
	s.Start()
}

