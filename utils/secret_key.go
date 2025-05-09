package utils

import (
	"os"
	"sync"
)

var secretKey []byte
var once sync.Once

func getSecretKey() []byte {
	once.Do(func() {
		secretKey = []byte(os.Getenv("SECRET_KEY"))
	})
	return secretKey
}
