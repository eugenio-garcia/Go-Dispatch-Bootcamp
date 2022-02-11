package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
	"github.com/gorilla/mux"
)

type usecase interface {
	GetAllPokemons() (model.Pokemons, error)
	GetPokemonById(id int) (*model.Pokemon, error)
	// GetPokemonByName(name string) (*model.Pokemon, error)
}

type pokemonController struct {
	usecase usecase
}

func New(uc usecase) pokemonController {
	return pokemonController{usecase: uc}
}

func (pc pokemonController) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	log.Printf("In controller GetAllPokemons")
	pokemons, err := pc.usecase.GetAllPokemons()
	if err != nil {
		log.Println("getting all pokemons from usecase")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "getting all pokemons error: %v\n", err)
		return
	}

	jsonData, err := json.Marshal(pokemons)

	if err != nil {
		log.Println("error marshalling pokemons")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "error marshalling pokemons %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (pc pokemonController) GetPokemonById(w http.ResponseWriter, r *http.Request) {
	log.Printf("In controller GetPokemonById")

	// extract the path parameters
	vars := mux.Vars(r)

	// convert the id param into an int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid id: %v", err)

		log.Fatalf("converting id param into an int: %v", err)
	}

	pokemon, err := pc.usecase.GetPokemonById(id)
	if err != nil {
		log.Println("getting pokemon from usecase")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "getting pokemon error: %v\n", err)
		return
	}

	jsonData, err := json.Marshal(pokemon)

	if err != nil {
		log.Println("error marshalling pokemons")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintf(w, "error marshalling pokemons %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
