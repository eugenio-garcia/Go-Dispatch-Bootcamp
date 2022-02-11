package main

import (
	"log"
	"net/http"

	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/controller"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/router"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/service"
	"github.com/eugenio-garcia/Go-Dispatch-Bootcamp/usecase"
)

func main() {

	pokemonService := service.New(nil)
	pokemonUsecase := usecase.New(pokemonService)
	pokemonController := controller.New(pokemonUsecase)
	r := router.Setup(pokemonController)

	log.Println("Listening on Port 8080")
	http.ListenAndServe("localhost:8080", r)

}
