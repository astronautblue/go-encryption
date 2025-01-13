// decrypt.go
package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// DecryptResult adalah hasil dari dekripsi yang berupa plaintext
type DecryptResult struct {
	Plaintext string `json:"plaintext"`
}

// Decrypt mendekripsi ciphertext dengan AES CBC mode dan mengembalikan hasil dekripsi
func Decrypt(key []byte, ivHex string, ciphertextHex string) (*DecryptResult, error) {
	// Pastikan panjang key yang digunakan adalah valid (16, 24, atau 32 byte untuk AES)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}

	// Mengonversi IV dan ciphertext dari hex ke []byte
	iv, err := hex.DecodeString(ivHex)
	if err != nil {
		return nil, err
	}

	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return nil, err
	}

	// Membuat cipher block dengan kunci
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Menggunakan AES CBC mode untuk dekripsi
	cbc := cipher.NewCBCDecrypter(block, iv)

	// Dekripsi ciphertext
	plaintext := make([]byte, len(ciphertext))
	cbc.CryptBlocks(plaintext, ciphertext)

	// Menghapus padding (padding menggunakan PKCS7, yang digunakan oleh CBC)
	plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]

	// Mengembalikan hasil dekripsi sebagai string
	return &DecryptResult{
		Plaintext: string(plaintext),
	}, nil
}
