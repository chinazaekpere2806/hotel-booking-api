package main

import (
	"fmt"
	"hotelapi/handlers"
	"hotelapi/middleware"
	"hotelapi/routes"
	"hotelapi/storage"
	"log"
	"net/http"
)

func main() {
	var err error

	handlers.Allbookings, err = storage.LoadBookings()
	if err != nil {
		log.Fatal(err)
	}
	handlers.AllRooms, err = storage.LoadRooms()
	if err != nil {
		log.Fatal(err)
	}
	handlers.Allguests, err = storage.LoadGuests()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	handleme := middleware.Logger(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handleme,
	}

	fmt.Println("Server running on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
