package Functionality

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Geolocation struct {
	Query   string  `json:"query"`      // Public IP
	Country string  `json:"country"`    // Country
	Region  string  `json:"regionName"` // Region
	City    string  `json:"city"`       // City
	Lat     float64 `json:"lat"`        // Latitude
	Lon     float64 `json:"lon"`        // Longitude
	ISP     string  `json:"isp"`        // Internet Service Provider
}

func GetGeolocation(ip string) (*Geolocation, error) {
	//using ip-api.com to obtain public ip of device
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var location Geolocation
	if err := json.Unmarshal(body, &location); err != nil {
		return nil, err
	}
	return &location, nil
}
