package usecase

import (
	"log"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
)

type PokemonService interface {
	GetAllPokemons() (model.Pokemons, error)
	GetPokemonById(id int) (*model.Pokemon, error)
	// GetPokemonByName(name string) (*model.Pokemon, error)
	LoadPokemonToCSV() (bool, error)
}

type PokemonUsecase struct {
	service PokemonService
}

func New(s PokemonService) PokemonUsecase {
	return PokemonUsecase{
		service: s,
	}
}

func (pu PokemonUsecase) GetAllPokemons() (model.Pokemons, error) {
	log.Printf("In usecase GetAllPokemons")
	pokemons, err := pu.service.GetAllPokemons()

	if err != nil {
		return nil, err
	}

	//redis to update cache

	//

	return pokemons, nil
}

func (pu PokemonUsecase) GetPokemonById(id int) (*model.Pokemon, error) {
	log.Printf("In usecase GetPokemonById")
	pokemon, err := pu.service.GetPokemonById(id)

	if err != nil {
		return nil, err
	}

	//redis to update cache

	//

	return pokemon, nil
}

func (pu PokemonUsecase) LoadPokemonToCSV() (bool, error) {
	log.Printf("In usecase LoadPokemonToCSV")
	loaded, err := pu.service.LoadPokemonToCSV()

	if err != nil {
		return loaded, err
	}

	return loaded, nil
}
