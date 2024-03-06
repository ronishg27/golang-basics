package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello from greet")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang Modules</h1>"))

}

// commands related to mod
// go mod init github.com/acer/golang/23mymodules

// go list
// go list all
// go list -m all
// go list -m -versions github.com/gorilla/mux
// go mod why github.com/gorilla/mux

// go mod graph
// go mod vendor
// go mod tidy
// go mod verify
// go mod download
// go mod edit -replace github.com/gorilla/mux=github.com/gorilla/mux@v1.8.0
// go mod edit -require github.com/gorilla/mux@v1.8.0
