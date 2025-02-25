package main

import (
	"log"
	"sync"

	"github.com/Irooniam/msg/socks"
)

func main() {
	log.Println("starting directory service...")
	router, err := socks.NewRouter("router")
	if err != nil {
		log.Fatal("cant create router ", err)
	}

	log.Println(router)
	if err := router.Bind(9321); err != nil {
		log.Fatal("problem binding router ", err)
	}

	var wg sync.WaitGroup

	go router.Run()
	wg.Add(1)

	wg.Wait()
}
