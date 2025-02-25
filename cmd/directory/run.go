package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Irooniam/msg/internal/socks"
	"github.com/joho/godotenv"
)

const ENV_DIR_IP = "ENV_DIR_IP"
const ENV_DIR_PORT = "ENV_DIR_PORT"

func main() {
	log.Println("starting directory service...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var dip string
	var dport int

	if dip = os.Getenv(ENV_DIR_IP); dip == "" {
		log.Fatalf("Expected ip for directory but env var is %s", dip)
	}

	//env var will come through as string
	sdport := os.Getenv(ENV_DIR_PORT)
	if sdport == "" {
		log.Fatalf("Expected port for directory but en var is %s", sdport)
	}
	dport, err = strconv.Atoi(sdport)
	if err != nil {
		log.Fatalf("Unable to convert directory port into int: %s", err)
	}

	router, err := socks.NewZRouter("router")
	if err != nil {
		log.Fatal("cant create router ", err)
	}

	if err := router.Bind(dip, dport); err != nil {
		log.Fatal("cant bind router ", err)
	}

	var wg sync.WaitGroup
	go router.Run()
	wg.Add(1)

	wg.Wait()
}
