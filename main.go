package main

import (
	"BlitzPubSec/Functionality"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//step 1: Get public IP
	publicIP, err := Functionality.GetPublicIP()
	if err != nil {
		fmt.Println("Error fetching public IP: ", err)
		return
	}
	fmt.Println("Public IP Address: ", publicIP)

	//step 2: obtain geolocation info
	geo, err := Functionality.GetGeolocation(publicIP)
	if err != nil {
		fmt.Println("Error fetching geolocation:", err)
		return
	}

	// Step 3: Display the geolocation data
	fmt.Printf("Geolocation for IP %s:\n", geo.Query)
	fmt.Printf("Country: %s\n", geo.Country)
	fmt.Printf("Region: %s\n", geo.Region)
	fmt.Printf("City: %s\n", geo.City)
	fmt.Printf("Latitude: %.6f\n", geo.Lat)
	fmt.Printf("Longitude: %.6f\n", geo.Lon)
	fmt.Printf("ISP: %s\n", geo.ISP)
}
