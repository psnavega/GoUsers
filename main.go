package main

import (
	"api/src/config"
	"api/src/db"
	"api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routers.Generate()
	config.Load()
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed connection")
	}
	fmt.Println(conn)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), r))
}
