package main

import (
	"strconv"
	"strings"

	"github.com/ccallazans/desafio-tunts/internal/helpers"
	"github.com/ccallazans/desafio-tunts/internal/models"
	"github.com/xuri/excelize/v2"
)

func (f *ExcelApp) setHeaderData() error {
	/*
		Insert data into the file header
	*/
	f.logger.Println("Insert header data")

	// Header data
	data := []*models.HeaderRow{
		{Sheet: "Sheet1", Cell: "A1", Text: "Countries List"},
		{Sheet: "Sheet1", Cell: "A2", Text: "Name"},
		{Sheet: "Sheet1", Cell: "B2", Text: "Capital"},
		{Sheet: "Sheet1", Cell: "C2", Text: "Area"},
		{Sheet: "Sheet1", Cell: "D2", Text: "Currencies"},
	}

	// Insert row data
	for _, row := range data {
		err := f.file.SetCellValue(row.Sheet, row.Cell, row.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ExcelApp) setRowData(data []*models.Country) error {

	f.logger.Println("Insert row data")

	for row, country := range data {
		rowString := strconv.Itoa(row + 3)
		err := insertRow(f.file, rowString, country)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertRow(file *excelize.File, row string, country *models.Country) error {
	/*
		Insert response data into rows
	*/

	// Insert country name
	err := file.SetCellValue("Sheet1", "A"+row, country.Name.Common)
	if err != nil {
		return err
	}

	// Insert country capital
	if len(country.Capital) == 0 {
		country.Capital = append(country.Capital, "-")
	}

	err = file.SetCellValue("Sheet1", "B"+row, country.Capital[0])
	if err != nil {
		return err
	}

	// Insert country area
	if &country.Area == nil {
		err = file.SetCellValue("Sheet1", "C"+row, "-")
		if err != nil {
			return err
		}

	} else {
		err = file.SetCellValue("Sheet1", "C"+row, country.Area)
		if err != nil {
			return err
		}
	}

	// Insert country currencies
	if len(country.Currencies) == 0 {
		country.Currencies = map[string]interface{}{"-": nil}
	}

	curSymbols := helpers.GetKeys(country.Currencies)
	value := strings.Join(curSymbols[:], ",")

	err = file.SetCellValue("Sheet1", "D"+row, value)
	if err != nil {
		return err
	}

	return nil
}

func (f *ExcelApp) setStyles() error {
	/*
		Format file text style
	*/

	f.logger.Println("Insert styles")

	err := styleAreaColumn(f.file)
	if err != nil {
		return err
	}

	err = setTitleStyle(f.file)
	if err != nil {
		return err
	}

	err = setSubtitleStyle(f.file)
	if err != nil {
		return err
	}

	return nil
}
