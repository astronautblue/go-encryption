// encrypt.go
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// Struct untuk hasil enkripsi
type EncryptedData struct {
	IV      string `json:"iv"`
	Content string `json:"content"`
}

// Encrypt mengenkripsi plaintext dengan AES CBC mode dan mengembalikan objek EncryptedData
func Encrypt(key []byte, plaintext []byte) (*EncryptedData, error) {
	// Pastikan panjang key yang digunakan adalah valid (16, 24, atau 32 byte untuk AES)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}

	// Membuat IV (Initialization Vector) acak
	iv := make([]byte, aes.BlockSize) // AES Block size adalah 16 byte
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Membuat cipher block dengan kunci
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Menyusun plaintext untuk panjang yang sesuai dengan blok AES (multiple dari BlockSize)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append(plaintext, make([]byte, padding)...)

	// Mengenkripsi data
	cbc := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedText))
	cbc.CryptBlocks(ciphertext, paddedText)

	// Mengonversi hasil enkripsi dan IV ke dalam format hex
	ivHex := hex.EncodeToString(iv)
	ciphertextHex := hex.EncodeToString(ciphertext)

	// Mengembalikan struct EncryptedData
	return &EncryptedData{
		IV:      ivHex,
		Content: ciphertextHex,
	}, nil
}
