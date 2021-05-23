package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	config "gotbotpoc/config"
)

// See alternate IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// func Encrypt(text []byte) ([]byte, error) {
func Encrypt(plainTextIncoming string) (cipherText string) {

	fmt.Println(plainTextIncoming)
	text := []byte(plainTextIncoming)
	fmt.Println(text)
	secretKey := config.Config("SECRET_KEY")
	fmt.Println(secretKey)
	fmt.Println(3)
	key := []byte(secretKey) // 32 bytes
	fmt.Println(key)
	fmt.Println(3)

	block, err := aes.NewCipher(key)
	fmt.Println(block)
	if err != nil {
	}

	fmt.Println("if err != nil {}")
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	fmt.Println(cfb)
	fmt.Println(ciphertext)
	stringCipherText := fmt.Sprintf("%0x", ciphertext)
	return string(stringCipherText)
}

func Decrypt(cypherTextIncoming string) (plainTextIncoming string) {

	text := []byte(cypherTextIncoming)
	secretKey := config.Config("SECRET_KEY")
	key := []byte(secretKey) // 32 bytes

	block, err := aes.NewCipher(key)
	if err != nil {
		//return nil, err
	}
	if len(text) < aes.BlockSize {
		//
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	plainText, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		//return nil, err
	}
	return string(plainText)
}
