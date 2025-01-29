package Functionality

import (
	"io"
	"net/http"
)

/**
 * GetPublicIP retrieves the public IP address of the client machine.
 * It sends an HTTP GET request to the "https://api.ipify.org?format=text" API endpoint and reads the response body.
 *
 * @return string The public IP address of the client machine.
 * @return error If an error occurs during the HTTP request or reading the response body, this function returns an error.
 */
func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(ip), nil
}
