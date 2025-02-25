package main

import (
	"log"
	"net/http"

	"github.com/Irooniam/msg/internal/rest"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting up iHTTP...")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Tried to read .env but got error:%s", err)
	}

	handlers := rest.DefaultH{}

	hroute := httprouter.New()
	hroute.GET("/", handlers.Index)

	log.Fatal(http.ListenAndServe(":8080", hroute))
}
