package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	secret := "this is secret"
	urlAddress := "https://oapi.dingtalk.com/robot/send?access_token=XXXXXXXXXXXX"
	
	timestamp := time.Now().UnixMilli()
	
	
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	fmt.Println(sign)


	UrlAddress := fmt.Sprintf("%s&timestamp=%d&sign=%s", urlAddress, timestamp, sign)
	fmt.Println(UrlAddress)

	client := http.Client{}

	r := make(map[string]interface{})
	r["msgtype"] = "text"
	r["text"] = map[string]string{
		"content": "我就是我, 是不一样的烟火",
	}

	jsonByte, _ := json.Marshal(r)
	requestBody := bytes.NewBuffer(jsonByte)
	response, err := client.Post(UrlAddress, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("%+v", response)
}
