package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Response struct {
	Status          string `json:"status"`
	TypeSource      string `json:"type_source"`
	UserID          string `json:"user_id"`
	Timestamp       int64  `json:"timestamp"`
	TypeDestination string `json:"type_destination"`
	Message         struct {
		Text string `json:"text"`
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"message"`
	BotID string `json:"bot_id"`
}

var err error

func main() {

	err = godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("please consider environment varibles: %s \n", err)
	}
	url := "https://chat-api.one.th/message/api/v1/push_message"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("to", os.Getenv("ONECHAT_ID"))
	_ = writer.WriteField("bot_id", os.Getenv("BOT_ID"))
	_ = writer.WriteField("type", "file")

	file, errFile := os.Open("/Users/noppong/Desktop/17892.jpg") // image path
	if errFile != nil {
		fmt.Printf("[Send Error]: %s \n", errFile.Error())
		return
	}
	defer file.Close()

	part, errFormFile := writer.CreateFormFile("file", filepath.Base("/Users/noppong/Desktop/17892.jpg")) // image path
	if errFormFile != nil {
		fmt.Println("Error CreateFormFile", errFormFile)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(errFormFile)
		return
	}

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	bearToken := "Bearer " + os.Getenv("CHAT_TOKEN")
	req.Header.Add("Authorization", bearToken)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	// String to JSON
	var resJson Response
	json.Unmarshal(body, &resJson)
	fmt.Println(resJson.Status)
}
