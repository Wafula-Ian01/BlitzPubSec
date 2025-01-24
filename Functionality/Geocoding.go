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

/**
 * GetLocation retrieves the location details from the Nominatim OpenStreetMap API using the provided latitude and longitude coordinates.
 *
 * @param lat float64 The latitude coordinate of the location.
 * @param lon float64 The longitude coordinate of the location.
 *
 * @return *Location A pointer to the Location struct containing the address display name, address details, structural type, and type.
 * @return *Address A pointer to the Address struct containing the street, neighbourhood, village, city, state, country, and country code.
 * @return error An error if there is an issue with the API call or JSON unmarshalling.
 */
func GetLocation(lat, lon float64) (*Location, *Address, error) {
	//making call to nominatim osm api
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var locality Location
	var addr Address
	if err := json.Unmarshal(body, &addr); err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(body, &locality); err != nil {
		return nil, nil, err
	}

	return &locality, &addr, nil
}
