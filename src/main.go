package main

import (
	"log"
	"net/http"

	"idotno.fr/echo/routes"
	"idotno.fr/echo/services"
	"idotno.fr/echo/utils"
)

func main() {
	// Connect to db
	dbConn := utils.GetDbConnection()
	services.SetDB(dbConn)

	// Router
	r := routes.CreateRouter()

	// Websockets
	go services.HandleWsMessages()

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
