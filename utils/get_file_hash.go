package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
)

func GetFileHash(fileM *multipart.FileHeader) (string, error) {

	file, err := fileM.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString, nil
}
