package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/health", handler)
	fmt.Println("server started...")
	http.ListenAndServe(":8080", nil)
}
