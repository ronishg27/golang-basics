package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ronishg27/mongodbapi/routes"
)

func main() {
	fmt.Println("MongoDB CRUD Application")
	fmt.Println("Server is getting started...")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000.")

}
