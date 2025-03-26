package main

import (
	"log"
	"sync"

	"github.com/Irooniam/msg/services"
)

const ENV_DIR_IP = "ENV_DIR_IP"
const ENV_DIR_PORT = "ENV_DIR_PORT"

func main() {
	log.Println("starting directory service...")

	directory, err := services.NewDirectory()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	go directory.Run()
	wg.Add(1)

	wg.Wait()
}
