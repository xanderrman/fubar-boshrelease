package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Not so fast, I'm still broken :(\n"))
	})
	port := "8080"
	if os.Getenv("FUBAR_PORT") != "" {
		port = os.Getenv("FUBAR_PORT")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
