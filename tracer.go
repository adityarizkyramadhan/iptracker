package iptracker

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

const (
	linkWeb = "https://iplogger.org/ip-tracker/?ip="
)

type IPInfo struct {
	IPAddress      string
	Continent      string
	Country        string
	City           string
	Latitude       string
	Longitude      string
	Accuracy       string
	StateRegion    string
	DistrictCounty string
	ZipCode        string
	TimeZone       string
	LocalTime      string
	ISPProvider    string
	Organization   string
	Connection     string
	EuropeanUnion  string
	WeatherStation string
}

func Trace(ip string) (*IPInfo, error) {
	client := new(http.Client)
	ipInfo := IPInfo{}
	req, err := http.NewRequest(http.MethodGet, linkWeb+ip, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".ip-info div").Each(func(i int, element *goquery.Selection) {
		key := element.Find("div:first-child").Text()
		value := element.Find("div:last-child").Text()
		switch key {
		case "IP address":
			ipInfo.IPAddress = value
		case "Continent":
			ipInfo.Continent = value
		case "Country by IP":
			ipInfo.Country = value
		case "City by IP":
			ipInfo.City = value
		case "Latitude":
			ipInfo.Latitude = value
		case "Longitude":
			ipInfo.Longitude = value
		case "Accuracy":
			ipInfo.Accuracy = value
		case "State / Region":
			ipInfo.StateRegion = value
		case "District / County":
			ipInfo.DistrictCounty = value
		case "Zip code":
			ipInfo.ZipCode = value
		case "TimeZone":
			ipInfo.TimeZone = value
		case "Local time":
			ipInfo.LocalTime = value
		case "ISP / Provider":
			ipInfo.ISPProvider = value
		case "Organization":
			ipInfo.Organization = value
		case "Connection":
			ipInfo.Connection = value
		case "European union":
			ipInfo.EuropeanUnion = value
		case "Weather station":
			ipInfo.WeatherStation = value
		}
	})
	return &ipInfo, nil
}
