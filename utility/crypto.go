package utility

import (
	"crypto/aes"
	"crypto/cipher"
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

// ECB mode implementation
type ecbEncrypter struct {
	b cipher.Block
}

func NewECBEncrypter(b cipher.Block) *ecbEncrypter {
	return &ecbEncrypter{b}
}

func (x *ecbEncrypter) Encrypt(dst, src []byte) error {
	if len(src)%x.b.BlockSize() != 0 {
		return fmt.Errorf("ecb encrypt: input not full blocks")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
	return nil
}

func ComputeKCV_CBC_AES(key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize) // All zero IV
	encrypter := cipher.NewCBCEncrypter(block, iv)

	plaintext := make([]byte, aes.BlockSize) // All zero plaintext
	ciphertext := make([]byte, aes.BlockSize)
	encrypter.CryptBlocks(ciphertext, plaintext)

	return fmt.Sprintf("%X", ciphertext[:3]), nil
}

func ComputeKCV_ECB_AES(key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, aes.BlockSize) // All zero plaintext
	ciphertext := make([]byte, aes.BlockSize)
	ecb := NewECBEncrypter(block)
	err = ecb.Encrypt(ciphertext, plaintext)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", ciphertext[:3]), nil
}
