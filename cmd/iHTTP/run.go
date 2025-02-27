package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/rest"
	"github.com/Irooniam/msg/internal/socks"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting up iHTTP...")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Tried to read .env but got error:%s", err)
	}

	if err := socks.ChkDealerConf(); err != nil {
		log.Fatal(err)
	}
	log.Println("Dealer socket configuration looks good...")

	dealer, err := socks.NewDealer(os.Getenv(conf.MSG_DEALER_ID))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created dealer socket with ID: %s", os.Getenv(conf.MSG_DEALER_ID))

	//already did config check so we know env vars set
	routerUri := fmt.Sprintf("%s:%s", os.Getenv(conf.MSG_DIR_HOST), os.Getenv(conf.MSG_DIR_PORT))
	if err := dealer.Connect(routerUri); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected dealer socket to directory: %s", routerUri)

	//get directory host/port to connect dealer t
	var dhost, dport string
	if dhost = os.Getenv(conf.MSG_DIR_HOST); dhost == "" {
		log.Fatalf("env var %s is not set", dhost)
	}

	if dport = os.Getenv(conf.MSG_DIR_PORT); dport == "" {
		log.Fatalf("env var %s is not set", dport)
	}

	if err := dealer.Connect(fmt.Sprintf("%s:%s", dhost, dport)); err != nil {
		log.Fatalf("Cant connect to dealer: %s", err)
	}

	handlers := rest.DefaultH{}

	hroute := httprouter.New()
	hroute.GET("/", handlers.Index)

	log.Fatal(http.ListenAndServe(":8080", hroute))
}
