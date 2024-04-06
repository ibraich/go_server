package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "wrong directory", http.StatusNotAcceptable)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error err:%v", err)
		return
	}
	fmt.Fprintf(w, "Post Request is sccesfull \n")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "adress: %s\n", adress)

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Print("Server starting at port:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
