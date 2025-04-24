package services

import (
	"backend/utility"
	"encoding/base64"
	"fmt"

	"github.com/sirupsen/logrus"
)

type CryptoService interface {
	GenerateKey(int, string) (string, error)
	GenerateKCV(string, string, string) (string, error)
}

type cryptoService struct {
	logger *logrus.Logger
}

func NewCryptoService(logger *logrus.Logger) *cryptoService {
	return &cryptoService{
		logger: logger,
	}
}

func (cs *cryptoService) GenerateKey(size int, prng string) (string, error) {
	key, err := utility.GenerateKey(size)
	if err != nil {
		cs.logger.WithFields(logrus.Fields{
			"error ": err.Error(),
		}).Info("parsing get parameters")
		return "", err
	}
	// base64Key := base64.StdEncoding.EncodeToString(key)
	// return base64Key, nil
	return fmt.Sprintf("%02x", key), nil
}

func (cs *cryptoService) GenerateKCV(key string, mode string, cipher string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("error decoding Base64:")
	}
	if cipher == "AES" {
		if mode == "CBC" {
			return utility.ComputeKCV_CBC_AES(decodedBytes)
		} else if mode == "ECB" {
			return utility.ComputeKCV_ECB_AES(decodedBytes)
		} else {
			return "", fmt.Errorf("unsupported cipher mode of operation")
		}
	} else if cipher == "DES" {
		return "", fmt.Errorf("not Implemented")
	} else {
		return "", fmt.Errorf("unsupported cipher ")
	}
}
