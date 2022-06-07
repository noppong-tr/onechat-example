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

func main() {

	url := "https://chat-api.one.th/message/api/v1/push_message"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("to", "Ucfd58148ae2c5ccaa033c3ab41cf3a29")
	_ = writer.WriteField("bot_id", "Bdfc6c2f2c043570aab59d19d407f2ec7")
	_ = writer.WriteField("type", "file")

	file, errFile := os.Open("/Users/noppong/Desktop/17892.jpg")
	if errFile != nil {
		fmt.Errorf("[Send Error]: %s", errFile.Error())
	}
	defer file.Close()

	part, errFormFile := writer.CreateFormFile("file", filepath.Base("/Users/noppong/Desktop/17892.jpg"))
	if errFormFile != nil {
		fmt.Println("Error CreateFormFile", errFormFile)
		return
	}

	_, err := io.Copy(part, file)
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
	req.Header.Add("Authorization", "Bearer Ad76c22030f0f574987fc28599219033798e6e80c50634949ac92fbe51e4bcc121dcada52fc754e738fdb88dee018b180")

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
	var resJson Response
	json.Unmarshal(body, &resJson)
	fmt.Println(resJson.Status)
}
