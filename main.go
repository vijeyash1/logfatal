package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "an introduction to log.Fatal")

}

func main() {

	http.HandleFunc("/", d)
	fmt.Println("serves on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
