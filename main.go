package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":5001", r))
}
