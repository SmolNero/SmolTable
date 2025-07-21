// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting SmolTable on http://localhost:8080")

	//Test route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, Welome to SmolTable 🎉)
	})

	err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: %v", err)
	}