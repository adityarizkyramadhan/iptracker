package main

import (
	"fmt"
	"github.com/adityarizkyramadhan/iptracker"
)

func main() {
	data, err := iptracker.Trace("103.121.18.64")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("IP Address:" + data.IPAddress)
	fmt.Println("Continent:" + data.Continent)
	fmt.Println("Country:" + data.Country)
	fmt.Println("City:" + data.City)
	fmt.Println("Latitude:" + data.Latitude)
	fmt.Println("Longitude:" + data.Longitude)
	fmt.Println("Accuracy:" + data.Accuracy)
}
