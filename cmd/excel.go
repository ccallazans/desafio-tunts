package main

import (
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

func setHeader(file *excelize.File) error {
	err := file.SetCellValue("Sheet1", "A1", "Countries List")
	if err != nil {
		log.Printf("Set 'Countries List' failed: %s", err)
		return err
	}

	err = file.SetCellValue("Sheet1", "A2", "Name")
	if err != nil {
		log.Printf("Set 'Name' failed: %s", err)
		return err
	}

	err = file.SetCellValue("Sheet1", "B2", "Capital")
	if err != nil {
		log.Printf("Set 'Capital' failed: %s", err)
		return err
	}

	err = file.SetCellValue("Sheet1", "C2", "Area")
	if err != nil {
		log.Printf("Set 'Area' failed: %s", err)
		return err
	}

	err = file.SetCellValue("Sheet1", "D2", "Currencies")
	if err != nil {
		log.Printf("Set 'Currencies' failed: %s", err)
		return err
	}

	return nil
}

func insertRow(file *excelize.File, row string, country Country) error {
	// Insert name
	err := file.SetCellValue("Sheet1", "A"+row, country.Name.Common)
	if err != nil {
		log.Printf("Set 'country.Name' failed: %s", err)
		return err
	}

	// Insert capital
	if len(country.Capital) == 0 {
		country.Capital = append(country.Capital, "-")
	} else {
		err = file.SetCellValue("Sheet1", "B"+row, country.Capital[0])
		if err != nil {
			log.Printf("Set 'country.Capital' failed: %s", err)
			return err
		}
	}

	// Insert area
	if &country.Area == nil {
		err = file.SetCellValue("Sheet1", "C"+row, "-")
		if err != nil {
			log.Printf("Set 'country.Area' failed: %s", err)
			return err
		}
	} else {
		err = file.SetCellValue("Sheet1", "C"+row, country.Area)
		if err != nil {
			log.Printf("Set 'country.Area' failed: %s", err)
			return err
		}
	}

	// Insert currencies
	if len(country.Currencies) == 0 {
		err = file.SetCellValue("Sheet1", "D"+row, "-")
		if err != nil {
			log.Printf("Set 'country.Currencies' failed: %s", err)
			return err
		}
	} else {
		curSymbols := getKeys(country.Currencies)
		value := strings.Join(curSymbols[:], ",")

		err = file.SetCellValue("Sheet1", "D"+row, value)
		if err != nil {
			log.Printf("Set 'country.Currencies' failed: %s", err)
			return err
		}
	}

	return nil
}

func getKeys(m map[string]interface{}) []string {

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
