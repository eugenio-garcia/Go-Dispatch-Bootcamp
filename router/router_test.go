package router

import (
	"net/http"
	"testing"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	type args struct {
		c controller
	}
	tests := []struct {
		name      string
		routeName string
		path      string
		methods   []string
	}{
		{
			name:      "GetAllPokemons route",
			routeName: "GetAllPokemons",
			path:      "/api/v1/pokemons",
			methods:   []string{http.MethodGet},
		},
		{
			name:      "GetPokemonById route",
			routeName: "GetPokemonById",
			path:      "/api/v1/pokemons/{id}",
			methods:   []string{http.MethodGet},
		},
		{
			name:      "LoadPokemonToCSV route",
			routeName: "LoadPokemonToCSV",
			path:      "/api/v1/load/pokemons",
			methods:   []string{http.MethodGet},
		},
	}

	r := Setup(&mocks.Controller{})

	for _, tt := range tests {

		// get the registered route
		route := r.Get(tt.routeName)

		if route == nil {
			t.Errorf("%s route is not registered - should be registered", tt.routeName)
			t.FailNow()
		}

		name := route.GetName()
		if name != tt.routeName {
			t.Errorf("route name is: %s - you have: %s", name, tt.routeName)
			t.FailNow()
		}

		path, _ := route.GetPathTemplate()
		if path != tt.path {
			t.Errorf("route path is: %s - you have: %s", path, tt.path)
			t.FailNow()
		}

		methods, _ := route.GetMethods()
		// deep equal check
		assert.EqualValues(t, methods, tt.methods, "route methods are not equal")
	}
}
