package pkg

import (
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetPath(name string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	_, err := os.Stat(dirPath)
	if !os.IsNotExist(err) {
		return dirPath, nil
	}

	return "", fmt.Errorf("folder with that name doesn't exist")
}

func GetPathToFile(name string, ext string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName+ext)

	_, err := os.Stat(dirPath)
	if !os.IsNotExist(err) {
		return dirPath, nil
	}

	return "", fmt.Errorf("folder with that name doesn't exist")
}

func GenerateUniqueFolder(name string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return dirPath, nil
	}

	return "", fmt.Errorf("directory already exists")
}

func GenerateUniqueFile(fileName string, name string, folder string, requiredExt string) (string, error) {
	ext := filepath.Ext(fileName)
	if ext != requiredExt {
		return "", fmt.Errorf("wrong file extension")
	}

	nameWithoutExt := name

	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(nameWithoutExt, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	filePath := filepath.Join(folder, uniqueName+ext)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return filePath, nil
	}

	return "", fmt.Errorf("file already exists")
}

func CreateFolder(dirPath string) error {
	if err := os.Mkdir(dirPath, fs.ModePerm); !os.IsNotExist(err) {
		return err
	}

	return nil
}

func CreateFile(file multipart.File, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	if err = out.Close(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	return nil
}
