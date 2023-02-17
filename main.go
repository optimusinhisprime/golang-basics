package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm failed: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name: %v Password: %v \n", name, password)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed.", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello")
}

func main() {
	//we will be using this to server or html pages to the client.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting the server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
