package main

import (
	"log"
	"sync"

	"github.com/Irooniam/msg/services/directory"
	"github.com/joho/godotenv"
)

const ENV_DIR_IP = "ENV_DIR_IP"
const ENV_DIR_PORT = "ENV_DIR_PORT"

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	log.Println("starting directory service...")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Tried to read .env but got error:%s", err)
	}

	directory, err := directory.New()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	go directory.RouterRun()
	go directory.RecvMsg()
	wg.Add(2)

	wg.Wait()
}
