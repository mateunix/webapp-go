package main

import (
	"fmt"
	"net/http"

	"github.com/mateunix/webapp-go/pkg/handlers"
)

const portNumber = ":8080"

// Func main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
