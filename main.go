package main

import (
	"fmt"
	"net/http"
)

var ports = [4]string{"8081", "8082", "8083", "8084"}

func main() {
	for _, port := range ports {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "root %s", port)
		})
		mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "test %s", port)
		})
		serverURI := fmt.Sprintf(":%s", port)
		go http.ListenAndServe(serverURI, mux)
	}
	select {}
}
