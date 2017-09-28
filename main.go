package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello friend!"))
}

func main() {
	fmt.Printf("Starting Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	fmt.Printf("Error Starting app", err , "\n")
}