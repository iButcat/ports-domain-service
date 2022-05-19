package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"port-domain-service/infrastructure/persistence"
	"port-domain-service/interfaces"
	"port-domain-service/services"
	"syscall"
)

func main() {
	// should be on a config struct initialized by a function that reads the environment variables.
	port := ":8080"

	portRepository := persistence.NewPortRepositoryImplt()
	portService := services.NewPortDomainService(portRepository)
	router := interfaces.RegisterRoutes(portService)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		// Wait here until CTRL-C or other term signal is received.
		log.Println("Server is now running. Press CTRL-C to exit.")
		errs <- http.ListenAndServe(port, router)
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
	}()

	log.Print("\nexit", <-errs)
}
