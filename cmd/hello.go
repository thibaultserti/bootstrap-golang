package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Serving on http://localhost:8080 ...")
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
