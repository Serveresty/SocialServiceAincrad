package utils

import (
	"SocialServiceAincrad/configs"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateFilename() (string, error) {
	lengthEnv := configs.GetEnv("FILENAME_SIZE")
	length, err := strconv.Atoi(lengthEnv)
	if err != nil {
		return "", err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(charset[r.Intn(len(charset))])
	}
	return builder.String(), nil
}
