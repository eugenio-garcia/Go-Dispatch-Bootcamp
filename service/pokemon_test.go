package service

import (
	"reflect"
	"testing"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
)

var (
	anyPokemonMap = PokemonMap{
		1: {
			ID:   1,
			Name: "John Doe",
		},
		2: {
			ID:   2,
			Name: "Jane Doe",
		},
		3: {
			ID:   3,
			Name: "John Smith",
		},
	}

	anyPokemons = model.Pokemons{
		{
			ID:   1,
			Name: "John Doe",
		},
		{
			ID:   2,
			Name: "Jane Doe",
		},
		{
			ID:   3,
			Name: "John Smith",
		},
	}
)

func TestNew(t *testing.T) {
	type args struct {
		pm PokemonMap
	}
	tests := []struct {
		name string
		pm   PokemonMap
		want PokemonService
	}{
		{
			name: "Happy path - sending an initialized map",
			pm:   anyPokemonMap,
			want: PokemonService{
				data: anyPokemonMap,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.pm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonService_GetAllPokemons(t *testing.T) {
	type fields struct {
		data PokemonMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Pokemons
		dbError error
	}{
		{
			name: "Happy path - should return all employees",
			fields: fields{
				data: anyPokemonMap,
			},
			want: anyPokemons,
		},
		// {
		// 	name:    "Should fail - data has not been initialized",
		// 	dbError: errors.ErrDataNotInitialized,
		// },
		// {
		// 	name:    "Should fail - data is empty",
		// 	fields:  fields{data: EmployeeMap{}},
		// 	dbError: errors.ErrEmptyData,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			es := PokemonService{
				data: tt.fields.data,
			}
			got, err := es.GetAllPokemons()

			if err != tt.dbError {
				t.Errorf("PokemonService.GetAllPokemons() error = %v, wantErr %v", err, tt.dbError)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonService.GetAllPokemons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonService_GetPokemonById(t *testing.T) {
	type fields struct {
		data PokemonMap
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Pokemon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PokemonService{
				data: tt.fields.data,
			}
			got, err := ps.GetPokemonById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonService.GetPokemonById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PokemonService.GetPokemonById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPokemonService_LoadPokemonToCSV(t *testing.T) {
	type fields struct {
		data PokemonMap
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PokemonService{
				data: tt.fields.data,
			}
			got, err := ps.LoadPokemonToCSV()
			if (err != nil) != tt.wantErr {
				t.Errorf("PokemonService.LoadPokemonToCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PokemonService.LoadPokemonToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
