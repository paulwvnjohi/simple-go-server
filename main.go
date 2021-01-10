package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reqUri := fmt.Sprintf(`
	RequestURI: %v 
	Host: %v
	Form: %v
	sorm: %v`,
		r.RequestURI, r.Host, r.Form, r.Form.Get("some"))
	log.Printf(reqUri)
	fmt.Fprintf(w, reqUri)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	wordingDir, err := os.Getwd()
	if err != nil {
		//handle error
		return
	}
	if r.Method == "POST" {
		serveForm(w, r)
		return
	}
	//path := filepath.Join(wordingDir, r.URL.Path)
	http.ServeFile(w, r, wordingDir+r.URL.Path)
}

func main() {
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/form", serveForm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
