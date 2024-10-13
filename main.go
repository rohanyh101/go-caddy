package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handled request from %s\n", r.RemoteAddr)
		w.Write([]byte("HI MOM!\n"))
	})

	log.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
