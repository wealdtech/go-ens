package ens

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ccipFetch(sender common.Address, data string, urls []string) (result string, err error) {
	for _, url := range urls {
		method := "POST"
		if strings.Contains(url, "{data}") {
			method = "GET"
		}

		body := []byte{}
		if method == "POST" {
			body, err = json.Marshal(map[string]interface{}{
				"data":   data,
				"sender": sender,
			})
			if err != nil {
				return "", err
			}
		}

		url = strings.ReplaceAll(url, "{sender}", sender.String())
		url = strings.ReplaceAll(url, "{data}", data)
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		var responseData map[string]interface{}
		var result string
		if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
			err = json.NewDecoder(resp.Body).Decode(&responseData)
			if err != nil {
				continue
			}
			var ok bool
			result, ok = responseData["data"].(string)
			if !ok {
				err = errors.New("invalid response from gateway")
				continue
			}
		} else {
			responseBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			result = string(responseBytes)
		}

		if resp.StatusCode != http.StatusOK {
			continue
		}

		return result, nil
	}

	return "", err
}
