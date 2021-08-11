package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, _ *http.Request) {
	message := []byte("Hello World")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
func defaultHandler(writer http.ResponseWriter, _ *http.Request) {
	html, err := template.ParseFiles("guestbook.html")
	if err != nil {
		log.Fatal(err)
	}

	guestBook := GuestBook{SignatureCount: 2, Signatures: []string {"I'm the first", "Had a great time"}}
	err = html.Execute(writer, guestBook)
	if err != nil {
		log.Fatal(err)
	}
}

func addSignatureHandler(writer http.ResponseWriter, _ *http.Request) {
	html, err := template.ParseFiles("signature.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(writer, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func saveSignatureHandler(writer http.ResponseWriter, req *http.Request) {
	signature := req.FormValue("Signature")
	fmt.Println(signature)

	http.Redirect(writer, req, "/", http.StatusFound)
}

type GuestBook struct {
	SignatureCount int
	Signatures []string
}

func main() {
	http.HandleFunc("/guestbook/new/", addSignatureHandler)
	http.HandleFunc("/guestbook/save/", saveSignatureHandler)
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
