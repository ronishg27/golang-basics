package routes

import (
	"github.com/gorilla/mux"
	"github.com/ronishg27/mongodbapi/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/delete-all-movies", controllers.DeleteAllMovie).Methods("DELETE")

	return r

}
