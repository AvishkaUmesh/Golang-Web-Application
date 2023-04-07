package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

// Home is a function to handle home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is Home Page")
}

// About is a function to handle about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(10, 20)
	fmt.Fprintf(w, fmt.Sprintf("Hello, World! This is About Page. Sum of 10 and 20 is %d", sum))
}

// addValues is a function to add two numbers
func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is running on port", PORT)
	http.ListenAndServe(PORT, nil)
}
