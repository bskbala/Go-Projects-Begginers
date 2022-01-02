package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	requestBody, error := json.Marshal(map[string]string{
		"username": "saikumar go",
		"email":    "balabsk58@gmial.com",
	})
	if error != nil {
		print(error)
	}

	response, error := http.Post("http://httpbin.org/post",
		"application/json", bytes.NewBuffer(requestBody))
	if error != nil {
		print(error)

	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		print(error)

	}
	fmt.Print(string(body))
}
