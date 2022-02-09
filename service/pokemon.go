package service

import (
	"log"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
)

type PokemonMap map[int]model.Pokemon

var pokemonsOrder []int = []int{1, 2}

var db PokemonMap = map[int]model.Pokemon{
	1: {ID: 1, Name: "Pikachu"},
	2: {ID: 2, Name: "Bulbasur"},
}

type PokemonService struct {
	data PokemonMap
}

func New(pm PokemonMap) *PokemonService {
	if pm == nil {
		pm = db
	}

	return &PokemonService{
		data: pm,
	}
}

func (ps *PokemonService) GetAllPokemons() (model.Pokemons, error) {
	log.Printf("In service GetAllPokemons")
	//some logic happens here
	//we obtain pokemons

	pokemons := make(model.Pokemons, 0, len(ps.data))

	for _, id := range pokemonsOrder {
		pokemons = append(pokemons, ps.data[id])
	}

	return pokemons, nil
}
