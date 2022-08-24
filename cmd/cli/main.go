package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type ExcelApp struct {
	logger *log.Logger
	file   *excelize.File
}

func NewExcelApp(logger *log.Logger) *ExcelApp {
	return &ExcelApp{
		logger: logger,
		file:   excelize.NewFile(),
	}
}

func main() {

	// Create logger
	logger := log.Default()

	// Make request to API
	data, err := getApiData(logger)
	if err != nil {
		logger.Fatal(err)
	}

	// Create AppExcel
	myApp := NewExcelApp(logger)

	// Define file header
	err = myApp.setHeaderData()
	if err != nil {
		myApp.logger.Fatal(err)
	}

	// Fill rows with response data
	err = myApp.setRowData(data)
	if err != nil {
		myApp.logger.Fatal(err)
	}

	// Format file with text styles
	err = myApp.setStyles()
	if err != nil {
		myApp.logger.Fatal(err)
	}

	// Export File
	myApp.logger.Println("Export file to result.xlsx")
	if err := myApp.file.SaveAs("result.xlsx"); err != nil {
		myApp.logger.Fatal(err)
	}
}
