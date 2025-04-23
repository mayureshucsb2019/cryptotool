package services

import (
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
	cs.logger.Infof("Generating key ... size %d; prng %s", size, prng)
	return "this is fake key", nil
}
