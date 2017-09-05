package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "http://127.0.0.1:8000"

func main() {
	var url string
	var err error
	var resp *http.Response

	key := time.Now().UnixNano()

	// clear cache
	url = fmt.Sprintf("%s/v1/cache/%d", baseURL, key)
	fmt.Printf("Clear Cache - %s\n", url)
	resp, err = http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)

	// set cache
	url = fmt.Sprintf("%s/v1/cache", baseURL)
	json := fmt.Sprintf(`{"key":"%d","value":"{\"hello\":\"world\"}","ttl":30}`, key)

	var payLoad = []byte(json)

	fmt.Printf("Set Cache - %s\n", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payLoad))

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	// get cache
	url = fmt.Sprintf("%s/v1/cache/%d", baseURL, key)
	fmt.Printf("Get Cache - %s\n", url)
	resp, err = http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)
}
