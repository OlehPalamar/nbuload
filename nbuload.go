package nbuload

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const cRESTBaseURL string = "http://bank.gov.ua/NBUStatService/v1/statdirectory/exchangenew?json"

type NBURates struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	Exchangedate string  `json:"exchangedate"`
}

// printData function print loaded rates as table
func PrintData(data []NBURates) {
	var count int = len(data)
	fmt.Println("Rows found = ", count)
	fmt.Println("********************* HSP NBU Loader LoadReates() ************************************")
	fmt.Println("| r030 | Name                                     |  cc  |     rate     |     date   |")
	fmt.Println("--------------------------------------------------------------------------------------")
	for i := 0; i < count; i++ {
		fmt.Printf("| %4d ", data[i].R030)
		fmt.Printf("| %-40s ", data[i].Txt)
		fmt.Printf("| %4s ", data[i].Cc)
		fmt.Printf("| %12.4f ", data[i].Rate)
		fmt.Printf("| %10s |\n", data[i].Exchangedate)
	}
	fmt.Println("--------------------------------------------------------------------------------------")
}

// LoadRates() function for get exchange rates from Natioanal Bank of Ukraine (current date)
func LoadRates() []NBURates {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, cRESTBaseURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	var data []NBURates
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

// LoadRates() function for get exchange rates from Natioanal Bank of Ukraine (period)
func LoadRatesPeriod(from, to time.Time) []NBURates {

	var data []NBURates
	var partialData []NBURates
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	x := from
	for x.Before(to) || x.Equal(to) {
		req, err := http.NewRequestWithContext(context.Background(),
			http.MethodGet, cRESTBaseURL+"&date="+x.Format("20060102"), nil)
		if err != nil {
			panic(err)
		}
		req.Header.Add("User-Agent", "Learning Go")
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("unexpected status: got %v", res.Status))
		}
		err = json.NewDecoder(res.Body).Decode(&partialData)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(partialData); i++ {
			if partialData[i].R030 > 0 {
				data = append(data, partialData[i])
			}
		}
		x = x.AddDate(0, 0, 1)
	}
	return data
}
