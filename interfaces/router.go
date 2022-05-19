package interfaces

import (
	"port-domain-service/services"

	"github.com/gorilla/mux"
)

type PortDomainAPI struct {
	service services.PortDomainService
}

func RegisterRoutes(service services.PortDomainService) *mux.Router {
	api := &PortDomainAPI{
		service: service,
	}

	router := mux.NewRouter()

	router.HandleFunc("/ports", api.createPorts).Methods("POST")
	router.HandleFunc("/ports", api.updatePorts).Methods("PUT")

	return router
}
