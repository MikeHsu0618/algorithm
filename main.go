package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Error    error
	Response *http.Response
}

func main() {
	for i := 0; i < 1000; i++ {
		go getM3u8()
	}
	time.Sleep(20 * time.Second)
}

func getM3u8() {
	url := "https://mediapull.aecricex.xyz/live/stream648158/index.m3u8"
	resp, _ := http.Get(url)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", body)
}
