package router

import (
	"backend/handler"
	"backend/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func CryptoRoutes(r *mux.Router, logger *logrus.Logger) {

	cs := services.NewCryptoService(logger)
	ch := handler.NewCryptoHandler(logger, cs)

	r.HandleFunc("/generateKey", ch.GenerateKey).Methods("POST")
	r.HandleFunc("/generateKCV", ch.GenerateKCV).Methods("POST")
	r.HandleFunc("/xorComponents", ch.GenerateKCV).Methods("GET")

}
