package model

type Pokemons []Pokemon

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
