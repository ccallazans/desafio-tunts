package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccallazans/desafio-tunts/internal/models"
)

func getApiData(logger *log.Logger) ([]*models.Country, error) {

	logger.Println("Make request to API")

	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response data
	logger.Println("Read response data")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Struct response data into CountryArray type
	countries := []*models.Country{}
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}
