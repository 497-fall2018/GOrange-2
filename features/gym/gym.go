package gym

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/497-fall2018/GOrange-2/internal/config"
)

type Gym struct {
	Open  bool `json:"open"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{gymID}", GetGym(configuration))
	router.Delete("/{gymID}", DeleteGym(configuration))
	router.Post("/", CreateGym(configuration))
	router.Get("/", GetAllGyms(configuration))
	return router
}

func GetGym(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gymID := chi.URLParam(r, "gymID")
		gyms := Gym{
			Open:  false,
			Title: "Hello world",
			Body:  gymID,
		}
		render.JSON(w, r, gyms) // A chi router helper for serializing and returning json
	}
}

func DeleteGym(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		gymID := chi.URLParam(r, "gymID")
		response := make(map[string]string)
		response["message"] = "Deleted Gym " + gymID
		render.JSON(w, r, response) // Return some demo response
	}
}


func CreateGym(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created Gym successfully"
		render.JSON(w, r, response) // Return some demo response
	}
}

func GetAllGyms(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		gyms := []Gym{
			{
				Open:  false,
				Title: "Hello world",
				Body:  "Heloo world from planet earth",
			},
		}
		render.JSON(w, r, gyms) // A chi router helper for serializing and returning json
	}
}