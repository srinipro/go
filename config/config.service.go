package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetConfig to make sure that was configured
func (Config) GetConfig() Config {

	fPath := buildAbsPath("./config/dev.json")
	openFile := openFile(fPath)
	bytes := readAllBytes(openFile)

	var props Config

	if err := json.Unmarshal(bytes, &props); err != nil {
		log.Fatal(err)
	}

	return props
}

func openFile(filePath string) *os.File {
	openedFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return openedFile
}

func readAllBytes(file *os.File) []byte {
	bytes, er := ioutil.ReadAll(file)

	if er != nil {
		log.Fatal("Error while read file:", er)
	}

	return bytes
}

// To build absolute path form base directory/drive
func buildAbsPath(relativePath string) string {
	AbsPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatal("Error while building Absolute path:", err)
	}
	return AbsPath
}
