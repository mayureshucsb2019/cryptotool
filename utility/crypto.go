package utility

import (
	"crypto/rand"
	"fmt"
)

// generateKey generates a cryptographic key of the given bit size.
// Valid sizes are 128, 192, or 256.
func GenerateKey(size int) ([]byte, error) {
	if size != 64 && size != 128 && size != 192 && size != 256 {
		return nil, fmt.Errorf("invalid key size: must be 64, 128, 192, or 256 bits")
	}

	bytes := size / 8
	key := make([]byte, bytes)

	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("error generating random bytes: %v", err.Error())
	}

	return key, nil
}
