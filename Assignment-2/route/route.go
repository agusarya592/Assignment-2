package route

import (
	"assignment2/controller"
	"assignment2/repository"
	"assignment2/service"
	"database/sql"

	"github.com/gorilla/mux"
)

//localhost:8080/api/orders

func Init(router *mux.Router, db *sql.DB) {
	webRouter := router.NewRoute().PathPrefix("/api").Subrouter()

	orderRepository := repository.ProvideRepository(db)
	orderService := service.ProvideService(orderRepository)
	orderController := controller.ProvideController(webRouter, orderService)

	orderController.InitController()
}
