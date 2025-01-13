# Enkripsi dan Dekripsi dengan AES dalam Go

Repo ini menunjukkan bagaimana melakukan enkripsi dan dekripsi pesan menggunakan AES dalam mode CBC (Cipher Block Chaining) di Golang. Kode ini dipisahkan menjadi beberapa file, di mana setiap bagian berfungsi secara terpisah: satu untuk enkripsi dan satu lagi untuk dekripsi. Hasil dari kedua operasi ini dikembalikan dalam bentuk objek (struct), yang menjadikan kode lebih terstruktur dan mudah dipahami.

## Struktur Proyek

Struktur folder proyek adalah sebagai berikut:

```go
/go-encryption 
├── encryption 
│ └── encrypt.go 
├── decryption 
│ └── decrypt.go 
├── main.go 
├── go.mod 
└── go.sum
```


- **`encryption/encrypt.go`**: Berisi fungsi untuk mengenkripsi data menggunakan AES dalam mode CBC.
- **`decryption/decrypt.go`**: Berisi fungsi untuk mendekripsi data yang sudah dienkripsi.
- **`main.go`**: File utama yang menghubungkan antara fungsi enkripsi dan dekripsi, serta menampilkan hasilnya.

---

## Penjelasan

### 1. Encrypt

Fungsi `Encrypt` di dalam file `encrypt.go` digunakan untuk mengenkripsi plaintext menggunakan AES dengan mode CBC. Hasil enkripsi (IV dan ciphertext) dikembalikan dalam bentuk objek `EncryptedData`.

#### Struktur `EncryptedData`

```go
type EncryptedData struct {
	IV      string `json:"iv"`
	Content string `json:"content"`
}
```
```
Catatan: 
IV: Initialization Vector, digunakan untuk meningkatkan keamanan enkripsi.
Content: Hasil ciphertext yang telah dienkripsi dalam format hexadecimal.
```

### Fungsi Encrypt

```go
func Encrypt(key []byte, plaintext []byte) (*EncryptedData, error)
```
Fungsi ini mengenkripsi plaintext dengan menggunakan AES CBC, menghasilkan IV dan ciphertext dalam format hexadecimal, kemudian mengembalikan objek EncryptedData.

Contoh:
```go
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type EncryptedData struct {
	IV      string `json:"iv"`
	Content string `json:"content"`
}

func Encrypt(key []byte, plaintext []byte) (*EncryptedData, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}

	iv := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append(plaintext, make([]byte, padding)...)

	cbc := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedText))
	cbc.CryptBlocks(ciphertext, paddedText)

	ivHex := hex.EncodeToString(iv)
	ciphertextHex := hex.EncodeToString(ciphertext)

	return &EncryptedData{
		IV:      ivHex,
		Content: ciphertextHex,
	}, nil
}

```

### 1. Decrypt

Fungsi Decrypt di dalam file decrypt.go digunakan untuk mendekripsi ciphertext yang telah dienkripsi menggunakan AES CBC. Hasil dekripsi dikembalikan dalam bentuk objek DecryptResult.

Struktur DecryptResult

```go
type DecryptResult struct {
	Plaintext string `json:"plaintext"`
}
```

Plaintext: Hasil dari proses dekripsi dalam bentuk teks yang dapat dibaca.

### Fungsi Decrypt

```go
func Decrypt(key []byte, ivHex string, ciphertextHex string) (*DecryptResult, error)
```

Fungsi ini mendekripsi ciphertext yang diberikan, dan mengembalikan plaintext dalam bentuk string setelah padding dihapus.

Contoh:

```go
package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

type DecryptResult struct {
	Plaintext string `json:"plaintext"`
}

func Decrypt(key []byte, ivHex string, ciphertextHex string) (*DecryptResult, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}

	iv, err := hex.DecodeString(ivHex)
	if err != nil {
		return nil, err
	}

	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	cbc.CryptBlocks(plaintext, ciphertext)

	// Menghapus padding
	plaintext = plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]

	return &DecryptResult{
		Plaintext: string(plaintext),
	}, nil
}
```

## Cara jalankan code ini

#### 1. Inisialisasi Modul Go
> Jalankan perintah berikut di terminal untuk membuat file go.mod:
```go
go mod init go-enkripsi<bisa juga disesuaikan dengan nama yang kamu inginkan>
```

#### 2. Jalankan file main.go
> Jalankan perintah berikut di terminal untuk membuat file go.mod:
```go
go run main.go

```

#### Output:
> Setelah menjalankan program, output akan terlihat seperti ini:
```go
Encrypted Data (JSON):
{
  "iv": "e4a5f4f59f9e9390c7b5d1ffab9a3c4d",
  "content": "aad0a57d417d8e424eaf98a30f49335c10cb040975adf6d30f5b72c63b370e0d"
}

Decrypted Text: This is a secret message.

```