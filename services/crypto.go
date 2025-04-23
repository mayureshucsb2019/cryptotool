package services

import (
	"backend/utility"
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

type CryptoServiceInterface interface {
	GenerateKey(int, string) (string, error)
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
	base64Key := base64.StdEncoding.EncodeToString(key)
	return base64Key, nil
}
