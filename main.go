package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	tb "gopkg.in/tucnak/telebot.v2"
)

var telegramToken string

func init() {
	telegramToken = os.Getenv("TELEGRAM_TOKEN_FIX_IMG")
}

func main() {

	bot, err := tb.NewBot(tb.Settings{
		Token:  telegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln(err)
	}

	bot.Handle(tb.OnPhoto, func(m *tb.Message) {
		b, err := fixSizeImage(getFile(m.Photo.FileID))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("ok fixed size image")
		ph := &tb.Photo{File: tb.FromReader(bytes.NewBuffer(b))}
		bot.Send(m.Sender, ph)
	})

	go bot.Start()
	// for heroku app 
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}

// GetFile
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
	telegramToken, fileID)
	b := get(url)
	log.Println("ok get file info")
	err := json.Unmarshal(b, &file)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("ok unmarshal file")
	url = fmt.Sprintf("https://api.telegram.org/file/bot%s/%s",
	telegramToken, file.Result.FilePath)
	b = get(url)
	log.Println("ok get file")
	return b
}

// get
func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("ok get request")
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("ok read request")
	return b
}

// name file
func fileName(path string) string {
	return filepath.Base(path)
}

// resize image to fix telegram profile image
func fixSizeImage(b []byte) ([]byte, error) {
	read := bytes.NewReader(b)
	img, f, err := image.Decode(read)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rec := img.Bounds()
	width := rec.Dx()
	height := rec.Dy()
	log.Printf("Resize Image %d X %d %s", width, height, f)
	if height > width {
		height = width
	} else if width > height {
		width = height
	}
	// Resize the image
	resizeImg := imaging.Fit(img, width, height, imaging.Lanczos)
	// Create a new black background image
	bgImage := imaging.New(width, height, color.Black)
	// paste the resized images into background image.
	img = imaging.PasteCenter(bgImage, resizeImg)
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	err = imaging.Encode(writer, img, imaging.JPEG)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return buf.Bytes(), nil
}
