package main

import (
	"log"
	"sync"

	"github.com/Irooniam/msg/internal/socks"
)

func main() {
	log.Println("starting directory service...")
	router, err := socks.NewZRouter("router")
	if err != nil {
		log.Fatal("cant create router ", err)
	}

	if err := router.Bind("0.0.0.0", 9876); err != nil {
		log.Fatal("cant bind router ", err)
	}

	var wg sync.WaitGroup
	go router.Run()
	wg.Add(1)

	wg.Wait()
}
