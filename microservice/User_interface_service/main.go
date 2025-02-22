package main

import (
	"fmt"
	"log"
	"net/http"
	"user_interface/router"
)

func main() {
	r := router.Router()
	fmt.Println("Running in port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
