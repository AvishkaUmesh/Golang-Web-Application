package main

import (
	"fmt"
	"net/http"

	"github.com/AvishkaUmesh/Golang-Web-Application/pkg/handlers"
)

const PORT = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server is running on port", PORT)
	http.ListenAndServe(PORT, nil)
}
