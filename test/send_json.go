package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:55555/request"
	fmt.Println("URL: ", url)

	r := make(map[string]string)
	r["usecase"] = "reid"
	r["uri"] = "10.10.10.1/live"
	r["device"] = "gpu"

	mJson, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(mJson))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Custom-Header", "request")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
