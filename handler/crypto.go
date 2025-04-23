package handler

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type CryptoHandlerInterface interface {
	GenerateKey(http.ResponseWriter, *http.Request)
}

type cryptoHandler struct {
	cs     services.CryptoServiceInterface
	logger *logrus.Logger
}

func NewCryptoHandler(logger *logrus.Logger, cs services.CryptoServiceInterface) *cryptoHandler {
	return &cryptoHandler{
		cs:     cs,
		logger: logger,
	}
}

func (ch *cryptoHandler) GenerateKey(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if size != 64 && size != 128 && size != 192 && size != 256 {
		http.Error(w, "size of key should be 64, 128, 192, 256 bits", http.StatusBadRequest)
		ch.logger.Info("Key size not a multiple of 64 bits")
		return
	}

	if err != nil {
		http.Error(w, "please check the format of parameter size", http.StatusBadRequest)
		ch.logger.WithFields(logrus.Fields{
			"error ": err.Error(),
		}).Info("parsing get parameters")
		return
	}

	key, err := ch.cs.GenerateKey(size, r.URL.Query().Get("prng"))
	if err != nil {
		http.Error(w, "error generating key", http.StatusInternalServerError)
		ch.logger.WithFields(logrus.Fields{
			"error ": err.Error(),
		}).Info("generating key")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GenerateKeyResp{Key: key})
}
