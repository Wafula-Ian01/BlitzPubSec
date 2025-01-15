package Functionality

import (
	"io"
	"net/http"
)

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(ip), nil
}
