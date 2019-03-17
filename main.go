package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!!!!!!!!!!")
}

func main() {

	http.HandleFunc("/", handle)

	log.Print("server start port 443")

	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
