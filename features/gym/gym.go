package gym

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Gym struct {
	Open  bool `json:"open"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{gymID}", GetGym)
	router.Delete("/{gymID}", DeleteGym)
	router.Post("/", CreateGym)
	router.Get("/", GetAllGyms)
	return router
}

func GetGym(w http.ResponseWriter, r *http.Request) {
	gymID := chi.URLParam(r, "gymID")
	gyms := Gym{
		Open:  false,
		Title: "Hello world",
		Body:  gymID,
	}
	render.JSON(w, r, gyms) // A chi router helper for serializing and returning json
}

func DeleteGym(w http.ResponseWriter, r *http.Request) {
	gymID := chi.URLParam(r, "gymID")
	response := make(map[string]string)
	response["message"] = "Deleted Gym"
	render.JSON(w, r, response) // Return some demo response
}

func CreateGym(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Gym successfully"
	render.JSON(w, r, response) // Return some demo response
}

func GetAllGyms(w http.ResponseWriter, r *http.Request) {
	gyms := []Gym{
		{
			Open:  false,
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, gyms) // A chi router helper for serializing and returning json
}