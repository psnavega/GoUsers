package main

import (
	"api/src/config"
	"api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routers.Generate()
	config.Load()

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), r))
}
