package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// 不可逆加密
func EncryptPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

// 加密后进行判断
func ComparePasswords(password, hashedPassword string) bool {
	return EncryptPassword(password) == hashedPassword
}

const encryptionKey = "ILikeYouEveryDayyaDyrevEuoYekiLI" // 32 bytes key
// 可逆加密
func EncryptPhoneNumber(phoneNumber string) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	plaintext := []byte(phoneNumber)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	nonce := ciphertext[:aes.BlockSize]

	for i := range nonce {
		nonce[i] = byte(i) // Simple nonce for demonstration purposes; consider using a random nonce
	}

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 可逆解密
func DecryptPhoneNumber(encryptedPhoneNumber string) (string, error) {
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPhoneNumber)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func ComparePhone(phoneNumber, encryptedPhoneNumber string) bool {
	data, err := EncryptPhoneNumber(phoneNumber)
	if err != nil {
		return false
	}
	if data == encryptedPhoneNumber {
		return true
	}
	return false
}
