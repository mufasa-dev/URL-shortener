package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sync"
	"url-shortener/config"
	"url-shortener/schemas"
)

var logger config.Logger

var (
	urlStore    = make(map[string]string)
	secretKey   = []byte("secretaeskey12345678901234567890")
	mu          sync.Mutex
	lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

var db = config.GetSQLite()

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")
	if originalUrl == "" {
		http.Error(w, "Param URL on query is needed", http.StatusBadRequest)
		return
	}

	encryptedUrl := encrypt(originalUrl)
	shortId := generateShortId()
	mu.Lock()
	urlStore[shortId] = encryptedUrl
	mu.Unlock()

	url := schemas.Url{
		OriginalUrl: originalUrl,
		ShortId:     shortId,
	}
	if err := db.Create(&url).Error; err != nil {
		logger.Errorf("error creating url: %v", err)
		http.Error(w, "Error creating url", http.StatusNotFound)
		return
	}

	shortUrl := fmt.Sprintf("http://localhost:8080/%s", shortId)

	fmt.Fprintf(w, "The shrtener URL is: %s", shortUrl)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortId := r.URL.Path[1:]

	var url schemas.Url
	result := db.Where("ShortUrl = ?", shortId).Find(&url)
	if result.Error != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func encrypt(originalUrl string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalUrl)
	clipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := clipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(clipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(clipherText)
}

func generateShortId() string {
	b := make([]rune, 6)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(lettersRune))))
		if err != nil {
			log.Fatal(err)
		}

		b[i] = lettersRune[num.Int64()]
	}

	return string(b)
}
