package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Irooniam/msg/conf"
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

	if err := socks.ChkRouterConf(); err != nil {
		log.Fatal(err)
	}

	router, err := socks.NewZRouter("router")
	if err != nil {
		log.Fatal("cant create router ", err)
	}

	rstr := fmt.Sprintf("%s:%s", os.Getenv(conf.MSG_DIR_HOST), os.Getenv(conf.MSG_DIR_PORT))
	if err := router.Bind(rstr); err != nil {
		log.Fatal("cant bind router ", err)
	}

	var wg sync.WaitGroup
	go router.Run()
	wg.Add(1)

	wg.Wait()
}
