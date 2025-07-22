// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting SmolTable on http://localhost:8080")	// Registers a rout** for root path

	//Test route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w," Welome to SmolTable 🎉")
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: %v", err)
	}
 }/