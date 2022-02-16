package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller interface {
	GetAllPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemonById(w http.ResponseWriter, r *http.Request)
	// GetPokemonByName(w http.ResponseWriter, r *http.Request)
	LoadPokemonToCSV(w http.ResponseWriter, r *http.Request)
}

func Setup(c controller) *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/pokemons", c.GetAllPokemons).Methods(http.MethodGet).Name("GetAllPokemons")
	v1.HandleFunc("/pokemons/{id}", c.GetPokemonById).Methods(http.MethodGet).Name("GetPokemonById")
	// v1.HandleFunc("/pokemonsByName/{name}", c.GetPokemonByName).Methods(http.MethodGet).Name("GetPokemonByName")

	//Load Endpoints
	v1.HandleFunc("/load/pokemons", c.LoadPokemonToCSV).Methods(http.MethodGet).Name("LoadPokemonToCSV")

	return r
}
