package main

/*
NOTES: use html/template to pass variables to html: https://stackoverflow.com/questions/27971240/how-to-pass-just-a-variable-not-a-struct-member-into-text-html-template-golan

	- start docker service with `sudo service docker start`
*/

import (
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/index.html")
}

func api_get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reponse!")
}

func api_post(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST request from %s: test=%s\n", r.RemoteAddr, r.PostFormValue("test"))
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))

	// Define routes
	r.HandleFunc("/", home)
	r.HandleFunc("/api", api_get).
		Methods("GET")
	r.HandleFunc("/api", api_post).
		Methods("POST")

	// Start the server
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
