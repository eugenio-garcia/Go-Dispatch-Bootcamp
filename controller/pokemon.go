package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
)

type usecase interface {
	GetAllPokemons() (model.Pokemons, error)
	// GetPokemonById(id int) (*model.Pokemon, error)
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
