package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

var DATE = time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)

type ExchangeRate struct {
	CurrencyIsoCode string // Này là kiểu int, string để test
	Usually         decimal.Decimal
	InHouse         decimal.Decimal
	Date            time.Time
}

func main() {
	isoCodes := []string{"USD", "USD", "CAD", "CAD", "YBM", "YBM"}
	rates := []string{"101", "102", "201", "202", "301", "302"}
	listDate := getDate("2023/8/1", "2023/8/3")

	ex := getExchangeRate(isoCodes, rates, listDate)

	// TEST KẾT QUẢ ...
	for _, rate := range ex {
		if rate.CurrencyIsoCode == "YBM" {
			fmt.Println(rate)
		}
		//fmt.Println(rate)
	}
}

// removeElementDuplicate ...
func removeElementDuplicate(input []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueSlice := make([]string, 0)

	for _, value := range input {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = true
			uniqueSlice = append(uniqueSlice, value)
		}
	}
	return uniqueSlice
}

// getDate ...
func getDate(fDate, tDate string) []time.Time {
	// Convert date string to date
	fromDate, _ := time.Parse("2006-01-02", formatDateString(strings.ReplaceAll(fDate, "/", "-")))
	toDate, _ := time.Parse("2006-01-02", formatDateString(strings.ReplaceAll(tDate, "/", "-")))

	// Get list date from fromDate to toDate
	dateList := make([]time.Time, 0)

	currentDate := fromDate
	for currentDate.Before(toDate) || currentDate.Equal(toDate) {
		dateList = append(dateList, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dateList
}

// getExchangeRate ...
func getExchangeRate(listIsoCode []string, listRate []string, listDate []time.Time) (result []ExchangeRate) {
	// Init map
	exchangeRateMap := make(map[int][]string, 0)

	// Add isoCode to map ->  map[ 1:[USD] 2:[USD] 3:[CAD] 4:[CAD] ]
	for i, iso := range listIsoCode {
		exchangeRateMap[i] = append(exchangeRateMap[i], iso)
	}

	// Add rate to map by key ->  map[ 1:[USD 101] 2:[USD 102] 3:[CAD 201] 4:[CAD 202] ]
	for i, value := range listRate {
		exchangeRateMap[i] = append(exchangeRateMap[i], value)

	}
	{
		fmt.Println(exchangeRateMap)
	}

	// -> map[ CAD:[201 202] USD:[101 102] ]
	outputMap := make(map[string][]string, 0)
	for _, value := range exchangeRateMap {
		currency := value[0]
		if _, ok := outputMap[currency]; !ok {
			outputMap[currency] = []string{}
		}

		outputMap[currency] = append(outputMap[currency], value[1])
	}

	// Map data vào struct model
	result = make([]ExchangeRate, 0)
	for _, date := range listDate {
		for key, value := range outputMap {
			ex := ExchangeRate{
				CurrencyIsoCode: key, // Cần get id by key (isoCode), string chỉ để test
				Date:            date,
			}

			if value[0] != "" && value[1] != "" {
				uData, _ := decimal.NewFromString(value[0])
				iData, _ := decimal.NewFromString(value[1])

				// Handle error

				ex.Usually = uData
				ex.InHouse = iData

			}

			result = append(result, ex)
		}
	}

	return
}

// formatDateString ...
func formatDateString(inputDate string) string {
	parsedTime, err := time.Parse("2006-1-2", inputDate)
	if err != nil {
		return ""
	}
	return parsedTime.Format("2006-01-02")
}

/*
[Taurus][TFN-1214]
- Upload file csv lên s3
- GetUrl download file về cho Client
*/
