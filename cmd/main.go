package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	countries := CountryArray{}
	err = json.Unmarshal(body, &countries)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	file := excelize.NewFile()
	err = setHeader(file)
	if err != nil {
		log.Fatalln(err)
	}

	for row, country := range countries {
		rowShift := row + 3
		rowString := strconv.Itoa(rowShift)
		insertRow(file, rowString, country)
	}
	// err = file.MergeCell("Sheet1", "A1", "A4")
	// if err != nil {
	// 	log.Printf("Merge cells failed: %s", err)
	// 	return
	// }
	if err := file.SaveAs("result.xlsx"); err != nil {
		fmt.Println(err)
	}

}
