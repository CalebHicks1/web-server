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

var working_dir, static_dir, port string
var local bool

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, working_dir+"static/index.html")
}

func api_get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reponse!")
}

func api_post(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST request from %s: test=%s\n", r.RemoteAddr, r.PostFormValue("test"))
}

func main() {

	local = false // set to true if running on local machine
	var srv *http.Server
	r := mux.NewRouter()

	if local == true {
		port = "8000"
		working_dir = "../"
		srv = &http.Server{
			Handler: r,
			Addr:    "localhost:8000",
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
	} else {
		port = "80"
		working_dir = "/usr/"
		srv = &http.Server{
			Handler: r,
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
	}

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(working_dir)))

	// Define routes
	r.HandleFunc("/", home)
	r.HandleFunc("/api", api_get).
		Methods("GET")
	r.HandleFunc("/api", api_post).
		Methods("POST")

	// Start the server

	log.Fatal(srv.ListenAndServe())
}
