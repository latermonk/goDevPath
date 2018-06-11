package main

import (
	"io"
	"log"
	"net/http"
)

func helloGoServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Hello, this is a Server")

}

func main() {
	http.HandleFunc("/", helloGoServer)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
