package main

import (
	"fmt"
	"net/http"

	"github.com/elue-dev/go-program/pkg/handlers"
)


var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Go Server running on PORT", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
