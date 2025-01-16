package Functionality

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name    string  `json:"display_name"` //Address Display name
	Address Address `json:"address"`      //Location Address details
	Class   string  `json:"class"`        //Structural type
	Type    string  `json:"type"`         //Type
}

type Address struct {
	Road          string `json:"road"`          // Street
	Neighbourhood string `json:"neighbourhood"` // Neighbourhood
	Village       string `json:"village"`       // Region
	City          string `json:"city"`          // City
	State         string `json:"state"`         // Region
	Country       string `json:"country"`       // Country
	Code          string `json:"country_code"`  // Country Code
}

func GetLocation(lat, lon float64) (*Location, error) {
	//making call to nominatim osm api
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var locality Location
	if err := json.Unmarshal(body, &locality); err != nil {
		return nil, err
	}

	return &locality, nil
}
