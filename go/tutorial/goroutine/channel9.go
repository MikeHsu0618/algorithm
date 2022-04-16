package main

import (
	"fmt"
	"net/http"
)

/*
	如何在 gorutine 中得知是哪一個 gorutine 出現錯誤
	通常會做出個 response struct 來裝載錯誤跟回傳訊息
*/

func gorutineErrorHandling() {

	type Response struct {
		Error    error
		Response *http.Response
	}

	checkUrl := func(urls []string) <-chan Response {

		respChan := make(chan Response)

		go func() {
			defer close(respChan)

			for _, url := range urls {
				var response Response
				resp, err := http.Get(url)
				response = Response{err, resp}
				respChan <- response
			}
		}()

		return respChan
	}

	respChan := checkUrl([]string{"https://google.com", "http://what"})
	for resp := range respChan {
		if resp.Error != nil {
			fmt.Println(resp.Error)
		} else {
			fmt.Println(resp.Response.StatusCode)
		}
	}
}
