package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/497-fall2018/GOrange-2/features/gym"
	"github.com/497-fall2018/GOrange-2/internal/config"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // sets content type to json
		middleware.Logger, // log api calls
		middleware.DefaultCompress, // compress results(mostly assets and json)
		middleware.RedirectSlashes, // redirect slashes to no slash url versions (basic parsing?)
		middleware.Recoverer, // recover from panic without crashing server
	)
	// version apis so you can update without breaking i
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/gym", gym.Routes(configuration))
	})

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := Routes(configuration)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Println("Serving application at PORT :" + configuration.Constants.PORT)
	log.Fatal(http.ListenAndServe(":8080", router)) // Note, the port is usually gotten from the environment.
}