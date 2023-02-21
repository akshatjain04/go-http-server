package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pong")); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/ping", ping)
	fmt.Println("Listening on port 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
