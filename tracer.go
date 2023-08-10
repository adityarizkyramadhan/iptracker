package iptracker

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

const (
	linkWeb = "https://iplogger.org/ip-tracker/?ip="
)

var (
	ErrGetData = errors.New("fail to get data")
)

type IPInfo struct {
	IPAddress string
	Continent string
	Country   string
	City      string
	Latitude  string
	Longitude string
	Accuracy  string
}

func Trace(ip string) (*IPInfo, error) {
	client := new(http.Client)
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
	if resp.StatusCode != http.StatusOK {
		fmt.Println(ErrGetData.Error())
		return nil, ErrGetData
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	ipInfo := IPInfo{
		IPAddress: doc.Find(".ip-info div:contains('IP address') .copy").AttrOr("data-copy", ""),
		Continent: doc.Find(".ip-info div:contains('Continent') .country").Text(),
		Country:   doc.Find(".ip-info div:contains('Country by IP') .country span").Text(),
		City:      doc.Find(".ip-info div:contains('City by IP') .country").Text(),
		Latitude:  doc.Find(".ip-info div:contains('Latitude') .icon span").Text(),
		Longitude: doc.Find(".ip-info div:contains('Longitude') .icon span").Text(),
		Accuracy:  doc.Find(".ip-info div:contains('Accuracy') span[data-named='content']").Text(),
	}

	return &ipInfo, nil
}
