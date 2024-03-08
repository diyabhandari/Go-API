package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" //used for web dev, other options available too
	//importing a package from our own module

	log "github.com/sirupsen/logrus" //log errors for debugging
	//can run go mod tidy to get these packages here in required in go.mod, but that didnt work so i used quick fix
)

// set up logger so that when we print something out we get the file and line number
func main() {
	log.SetReportCaller(true)        //pass true to turn this on
	var r *chi.Mux = chi.NewRouter() //function returns pointer to mux type, stored in r, its a struct we'll use to setup our API
	handler.Handlers(r)
	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost: 8000", r) //start the server with http package. it takes base location of the server and a handler
	if err != nil {
		log.Error(err) //log any errors we might have while opening the server
	}
}
