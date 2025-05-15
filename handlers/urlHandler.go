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
	"url-shortener/schemas"

	"github.com/gin-gonic/gin"
)

var (
	urlStore    = make(map[string]string)
	secretKey   = []byte("secretaeskey12345678901234567890")
	mu          sync.Mutex
	lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// @Sumary Shorten URL
// @Description Recive a URL and return a shorten version
// @Tags Shorten
// @Param url query string true "URL"
// @Success 200 {object}
// @Failure 401 {object}
// @Failure 500 {object}
// @Router /shorten [get]
func ShortenUrl(ctx *gin.Context) {
	originalUrl := ctx.Query("url")
	if originalUrl == "" {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Param URL on query is needed",
			"code":    http.StatusBadRequest,
		})
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
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error creating url",
			"code":    http.StatusNotFound,
		})
		return
	}

	shortUrl := fmt.Sprintf("http://localhost:8080/%s", shortId)

	fmt.Fprintf(ctx.Writer, "The shrtener URL is: %s", shortUrl)
}

func RedirectHandler(ctx *gin.Context) {
	shortId := ctx.Param(":id")

	var url schemas.Url
	result := db.Where("ShortUrl = ?", shortId).Find(&url)
	if result.Error != nil {
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "URL not found",
			"code":    http.StatusNotFound,
		})
		return
	}

	ctx.Redirect(http.StatusFound, url.OriginalUrl)
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
