package httpreq

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Call will make request to api and return []byte of response or error
func Call(url string, method string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	byteRes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return byteRes, nil
}