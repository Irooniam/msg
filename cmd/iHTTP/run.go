package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/rest"
	"github.com/Irooniam/msg/internal/socks"
	"github.com/joho/godotenv"
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

	if err := rest.ChkiHTTPConf(); err != nil {
		log.Fatal(err)
	}

	serverstr := fmt.Sprintf("%s:%s", os.Getenv(conf.MSG_IHTTP_HOST), os.Getenv(conf.MSG_IHTTP_PORT))
	server := rest.NewREST(serverstr, dealer)
	server.SetupHandlers()

	go dealer.Run()

	time.Sleep(time.Second * 1)

	for i := 0; i < 4; i++ {
		log.Println("trying to send from ihttp to router")
		dealer.SendMsg([]byte("router"))
		log.Println("ihttp sent", i)
	}

	log.Printf("Starting iHTTP listening on %s", serverstr)

	err = server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
