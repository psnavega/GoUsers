package main

import (
	"api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routers.Generate()

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":5001", r))
}
