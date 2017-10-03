package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/plain"}
	w.Header()["Cache-Control"] = []string{"no-cache"}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	w.Write([]byte(time.Now().String()))
}
