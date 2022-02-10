package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/errors"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
)

type PokemonMap map[int]model.Pokemon

var pokemonsOrder []int = []int{1, 2, 3}

var db PokemonMap = map[int]model.Pokemon{
	1: {ID: 1, Name: "Pikachu"},
	2: {ID: 2, Name: "Bulbasur"},
}

type PokemonService struct {
	data PokemonMap
}

func New(pm PokemonMap) *PokemonService {
	if pm == nil {
		pm = OpenCSV()
	}

	return &PokemonService{
		data: pm,
	}
}

type PokemonRecord struct {
	ID   int
	Name string
}

func createPokemonMap(data [][]string) map[int]model.Pokemon {
	var pokemonList PokemonMap = map[int]model.Pokemon{}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec PokemonRecord
			for j, field := range line {
				if j == 0 {
					rec.ID, _ = strconv.Atoi(field) //Error?
				} else if j == 1 {
					rec.Name = field
				}
			}
			pokemon := model.Pokemon{
				ID:   rec.ID,
				Name: rec.Name,
			}
			pokemonList[i] = pokemon
		}
	}
	return pokemonList
}

func OpenCSV() map[int]model.Pokemon {
	// open file
	f, err := os.Open("data/pokemons.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	pokemonList := createPokemonMap(data)

	// print the array
	fmt.Printf("%+v\n", pokemonList)
	return pokemonList

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

func (ps *PokemonService) GetPokemonById(id int) (*model.Pokemon, error) {
	log.Printf("In service GetPokemonById")

	// find the employee in the data
	pokemon, ok := ps.data[id]
	if !ok {
		//return nil, errz.ErrNotFound
		return nil, errors.ErrNotFound
	}

	return &pokemon, nil
}
