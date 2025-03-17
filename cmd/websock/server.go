package main

import (
	"log"

	"github.com/Irooniam/msg/internal/socks"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting websocket service...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := socks.ChkRouterConf(); err != nil {
		log.Fatal(err)
	}

}
