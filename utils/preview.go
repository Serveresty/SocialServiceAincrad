package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CreatePreview(videoPath string) (string, error) {
	absVideoPath, err := filepath.Abs(videoPath)
	if err != nil {
		fmt.Println("Error converting video path to absolute: " + err.Error())
		return "", err
	}

	videoDir := filepath.Dir(videoPath)
	videoFilename := filepath.Base(absVideoPath)
	previewDir := filepath.Join(videoDir, "previews")

	if _, err := os.Stat(previewDir); os.IsNotExist(err) {
		err = os.MkdirAll(previewDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating previews directory: " + err.Error())
			return "", err
		}
	}

	previewPath := filepath.Join(previewDir, fmt.Sprintf("%s.jpg", videoFilename))

	ffmpegPath := "C:\\ProgramData\\chocolatey\\bin\\ffmpeg.exe"

	cmd := exec.Command(ffmpegPath, "-i", videoPath, "-ss", "00:00:01.000", "-vframes", "1", previewPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running ffmpeg: " + err.Error())
		return "", err
	}
	return previewPath, nil
}

func ConvertToBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	base64Encoding := base64.StdEncoding.EncodeToString(fileBytes)
	return base64Encoding, nil
}
