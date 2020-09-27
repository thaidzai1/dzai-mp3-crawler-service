package httpreq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Call will make request to api and return []byte of response or error
func Call(url string, method string, body interface{}, resp chan []byte, errChan chan error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		errChan <-err
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		errChan <-err
		return
	}
	
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		errChan <-err
		fmt.Println("error: ", err, url)
		return
	}

	byteRes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errChan <-err
		fmt.Println("error read: ", err, url)
		return
	}

	resp <-byteRes
}