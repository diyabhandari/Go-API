package handlers

import (
	"github.com/diyabhandari/Go-API/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware" //middleware is the function that gets called before the primary function which handles the end point
)

func Handler(r *chi.Mux) { //title case helps the compiler understand that this function can be imported, if it was lower case itd have been private
	//add a piece of global middleware. its pre built, and taken from chi. global means itll be applied to all the end points
	r.Use(chimiddle.StripSlashes)                 //if the last / isnt slashed itll lead to a 404 error
	r.Route("/account", func(router chi.Router) { //set up a route, takes in a path and an anonymous func that takes chi.Router as parameter
		//we can use the router to define our get() function
		//adding another piece of middleware to check if user is authorized to access this data
		router.Use(middleware.Authorization) //2nd one is the func name
		//every request thsat wants to access an endpoint that starts with /account must pass this function first, if it fails and error is returned and the rest of the code isnt executed
		router.Get("/coins", GetCoinBalance) //2nd one is the func name

	})

}
