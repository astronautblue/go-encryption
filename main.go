// main.go
package main

import (
	"encoding/json"
	"fmt"
	"go-enkripsi/aes-256/decryption"
	"go-enkripsi/aes-256/encryption"
	"log"
)

func main() {
	// Kunci untuk enkripsi dan dekripsi (32 byte untuk AES-256)
	key := []byte("thisis32bitlongpassphraseimusing")

	// Pesan yang akan dienkripsi
	plaintext := []byte("This is a secret message.")

	// Enkripsi pesan
	encryptedData, err := encryption.Encrypt(key, plaintext)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	// Menampilkan hasil enkripsi dalam bentuk JSON
	encryptedJSON, _ := json.MarshalIndent(encryptedData, "", "  ")
	fmt.Println("Encrypted Data (JSON):")
	fmt.Println(string(encryptedJSON))

	// Dekripsi pesan
	decryptedData, err := decryption.Decrypt(key, encryptedData.IV, encryptedData.Content)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	// Menampilkan hasil dekripsi
	fmt.Println("\nDecrypted Text:", decryptedData.Plaintext)
}
