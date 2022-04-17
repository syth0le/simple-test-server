package main

import (
	"fmt"
	"net/http"
)

var ports = [4]string{"8081", "8082", "8083", "8084"}

func main() {
	//for _, port := range ports {
	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, "root %s", port)
	//	})
	//	serverURI := fmt.Sprintf(":%s", port)
	//	go http.ListenAndServe(serverURI, mux)
	//}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root %s", ":8081")
	})
	go http.ListenAndServe(":8081", mux)
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root %s", ":8082")
	})
	go http.ListenAndServe(":8082", mux2)
	mux3 := http.NewServeMux()
	mux3.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root %s", ":8083")
	})
	go http.ListenAndServe(":8083", mux3)
	mux4 := http.NewServeMux()
	mux4.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root %s", ":8084")
	})
	go http.ListenAndServe(":8084", mux4)

	select {}
}
