package main

import (
	"log"

	"github.com/Irooniam/msg/socks"
)

func main() {
	log.Println("starting directory service...")
	router, err := socks.NewRouter("router")
	log.Println(router, err)
}
