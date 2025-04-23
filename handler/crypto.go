package handler

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"

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
	// size, err := strconv.Atoi(r.URL.Query().Get("size"))
	body := models.GenerateKeyReq{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if body.Size != 64 && body.Size != 128 && body.Size != 192 && body.Size != 256 {
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

	key, err := ch.cs.GenerateKey(body.Size, body.PRNG)
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

func (ch *cryptoHandler) GenerateKCV(w http.ResponseWriter, r *http.Request) {
	// size, err := strconv.Atoi(r.URL.Query().Get("size"))
	body := models.GenerateKCVReq{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		ch.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Info("error decoding request body")
		return
	}

	if body.Mode != "CBC" && body.Mode != "ECB" {
		http.Error(w, "modes supported are CBC and ECB", http.StatusBadRequest)
		ch.logger.Info("unsupported mode of operation")
		return
	}

	if body.Cipher != "AES" && body.Cipher != "DES" {
		http.Error(w, "ciphersuite supported are AES and DES", http.StatusBadRequest)
		ch.logger.Info("unsupported cipher")
		return
	}
	// GenerateKCV(key string, mode string, cipher string) (string, error)
	kcv, err := ch.cs.GenerateKCV(body.Key, body.Mode, body.Cipher)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		ch.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Info("error generating KCV")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GenerateKCVResp{KCV: kcv})
}
