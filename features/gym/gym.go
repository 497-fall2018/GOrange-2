package gym

import (
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/497-fall2018/GOrange-2/internal/config"
)

type Gym struct {
	Open  string `json:"open"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{gymID}", GetGym(configuration))
	router.Delete("/{gymID}", DeleteGym(configuration))
	router.Post("/edit", EditGym(configuration))
	router.Post("/", CreateGym(configuration))
	router.Get("/", GetAllGyms(configuration))
	return router
}

func GetGym(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gymID := chi.URLParam(r, "gymID")
		gyms := Gym{
			Open:  "Full",
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

func EditGym(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var gym Gym
		if err := json.NewDecoder(r.Body).Decode(&gym); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}


		if err := configuration.Database.C("gym").Update(bson.M{"title" : gym.Title}, &gym); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		render.JSON(w, r, gym)
	}
}


func CreateGym(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var gym Gym
		if err := json.NewDecoder(r.Body).Decode(&gym); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := configuration.Database.C("gym").Insert(&gym); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		render.JSON(w, r, gym) // Return some demo response
	}
}

func GetAllGyms(configuration *config.Config) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		var gyms []Gym
		err := configuration.Database.C("gym").Find(bson.M{}).All(&gyms)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		render.JSON(w, r, gyms) // A chi router helper for serializing and returning json
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}