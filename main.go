package main

import (
	"first/controller"
	"first/model"
	"log"
	"net/http"
	"os"
)

func main() {
	model.Connect()
	mux := controller.Register()
	port := os.Getenv("PORT")
	log.Println("PORT: " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
