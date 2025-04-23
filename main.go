package main

import (
	"backend/router"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	logger := logrus.New()
	router.CryptoRoutes(r, logger)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST"},
		Debug:            true,
	})

	corsHandler := c.Handler(r)
	logger.Info("Starting server at port 8080")
	logger.Fatalf(http.ListenAndServe(":8080", corsHandler).Error())
}
