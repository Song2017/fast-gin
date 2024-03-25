package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Decrypt(encrypted string, key []byte) string {
	ciphertext, _ := DecryptBytes(encrypted, key)
	return string(ciphertext)
}

func DecryptBytes(encrypted string, key []byte) ([]byte, error) {
	ciphertext, _ := base64.RawURLEncoding.DecodeString(encrypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func Encrypt(idata string, key []byte) string {
	data := []byte(idata)
	encryptData, _ := EncryptBytes(data, key)
	return encryptData
}

func EncryptBytes(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

// base64
func B64Encode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

func B64Decode[T comparable](data T) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", data))
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
func B64DecodeBytes(data []byte) ([]byte, error) {
	var decoded []byte
	_, err := base64.StdEncoding.Decode(decoded, data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
