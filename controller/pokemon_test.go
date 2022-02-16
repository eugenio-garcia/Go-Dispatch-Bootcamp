package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/errors"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/mocks"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
	"github.com/gorilla/mux"
)

func TestNew(t *testing.T) {
	type args struct {
		uc usecase
	}
	tests := []struct {
		name string
		args args
		want pokemonController
	}{
		{
			name: "Happy path - create a new controller",
			args: args{
				uc: &mocks.Usecase{},
			},
			want: pokemonController{
				usecase: &mocks.Usecase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pokemonController_GetAllPokemons(t *testing.T) {
	type fields struct {
		usecase usecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := pokemonController{
				usecase: tt.fields.usecase,
			}
			pc.GetAllPokemons(tt.args.w, tt.args.r)
		})
	}
}

var (
	anyPokemon = &model.Pokemon{
		ID:   1,
		Name: "John Doe",
	}

	anyPathParams = map[string]string{
		"id": "1",
	}

	anyID = 1
)

func Test_pokemonController_GetPokemonById(t *testing.T) {
	tests := []struct {
		id          int
		name        string
		pathParams  map[string]string
		callUsecase bool
		ucPokemon   *model.Pokemon
		ucError     error
		want        int
	}{
		{
			name:        "Happy path - get a pokemon by id",
			id:          1,
			pathParams:  anyPathParams,
			callUsecase: true,
			ucPokemon:   anyPokemon,
			want:        http.StatusOK,
		},
		{
			name: "Should fail - no ID in the path",
			want: http.StatusBadRequest,
		},
		{
			name:        "Should fail - error not found while getting pokemon",
			id:          anyID,
			pathParams:  anyPathParams,
			callUsecase: true,
			ucError:     errors.ErrNotFound,
			want:        http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		uc := &mocks.Usecase{}

		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/pokemons/:id", nil)
			req = mux.SetURLVars(req, tt.pathParams)

			if tt.callUsecase {
				uc.On("GetPokemonById", tt.id).Return(tt.ucPokemon, tt.ucError)
			}

			pc := pokemonController{
				usecase: uc,
			}
			pc.GetPokemonById(rw, req)

			if rw.Code != tt.want {
				t.Errorf("pokemonController.GetPokemonById() = %v, want %v", rw.Code, tt.want)
			}

			uc.AssertExpectations(t)
		})
	}
}

func Test_pokemonController_LoadPokemonToCSV(t *testing.T) {
	type fields struct {
		usecase usecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := pokemonController{
				usecase: tt.fields.usecase,
			}
			pc.LoadPokemonToCSV(tt.args.w, tt.args.r)
		})
	}
}
