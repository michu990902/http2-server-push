package main

import (
	"log"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", Home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	
	log.Fatal(http.ListenAndServeTLS(":5000", "tls/cert_test_example.pem", "tls/key_test_example.pem", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	//server push
	//https://blog.golang.org/h2push
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push("/assets/script.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
		if err := pusher.Push("/assets/style.css", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	//
}