// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	model "github.com/eugenio-garcia/Go-Dispatch-Bootcamp/model"
	mock "github.com/stretchr/testify/mock"
)

// PokemonService is an autogenerated mock type for the PokemonService type
type PokemonService struct {
	mock.Mock
}

// GetAllPokemons provides a mock function with given fields:
func (_m *PokemonService) GetAllPokemons() (model.Pokemons, error) {
	ret := _m.Called()

	var r0 model.Pokemons
	if rf, ok := ret.Get(0).(func() model.Pokemons); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Pokemons)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPokemonById provides a mock function with given fields: id
func (_m *PokemonService) GetPokemonById(id int) (*model.Pokemon, error) {
	ret := _m.Called(id)

	var r0 *model.Pokemon
	if rf, ok := ret.Get(0).(func(int) *model.Pokemon); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadPokemonToCSV provides a mock function with given fields:
func (_m *PokemonService) LoadPokemonToCSV() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
