package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

func getFile(fileID string) []byte {

	var file struct {
		Ok     bool `json:"ok"`
		Result struct {
			FileID   string `json:"file_id"`
			FileSize uint32 `json:"file_size"`
			FilePath string `json:"file_path"`
		} `json:"result"`
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s",
		token, fileID)
	b,err := getRequest(url)
	log.Println("ok get file info")
	err = json.Unmarshal(b, &file)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("ok unmarshal file")
	url = fmt.Sprintf("https://api.telegram.org/file/bot%s/%s",
		token, file.Result.FilePath)
	b,err = getRequest(url)
	log.Println("ok get file")
	
	return b
}

// name file
func fileName(path string) string {
	return filepath.Base(path)
}
