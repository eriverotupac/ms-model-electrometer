package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"ms-model-electrometer/internal/config"
)

type Cipher struct {
	log          *zap.SugaredLogger
	keyForCipher string
}

func NewCipher(logger *zap.SugaredLogger, config config.Environment) *Cipher {
	return &Cipher{
		log:          logger,
		keyForCipher: config.KeyForCipher,
	}
}

func (c *Cipher) DecryptString(cipheredText string) (string, error) {
	keyHex := c.keyForCipher

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return "", err
	}

	decryptedText, err := executeDecryption(cipheredText, key)
	if err != nil {
		c.log.Error("Error in decryption:", err)
		return "", err
	}

	return decryptedText, nil
}

func executeDecryption(encryptedText string, key []byte) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCBCDecrypter(block, iv)
	cfb.CryptBlocks(cipherText, cipherText)

	return string(cipherText), nil
}
