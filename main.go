package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "Текущее время сервера: %s", currentTime)
	})

	http.ListenAndServe(":8080", nil)
}
