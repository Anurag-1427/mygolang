package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anurag-1427/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")

}
