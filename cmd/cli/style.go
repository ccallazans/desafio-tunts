package main

import "github.com/xuri/excelize/v2"


func setTitleStyle(file *excelize.File) error {
	/*
		Set text title format
	*/

	// Merge title cells
	err := file.MergeCell("Sheet1", "A1", "D1")
	if err != nil {
		return err
	}

	// Define cell style
	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#4F4F4F",
			Size:  16,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return err
	}

	// Apply modifications
	err = file.SetCellStyle("Sheet1", "A1", "D1", style)
	if err != nil {
		return err
	}

	return nil
}

func setSubtitleStyle(file *excelize.File) error {
	/*
		Set text subtitle format
	*/

	// Define cell style
	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#808080",
			Size:  12,
		},
	})
	if err != nil {
		return err
	}

	// Apply modifications
	err = file.SetCellStyle("Sheet1", "A2", "D2", style)
	if err != nil {
		return err
	}

	return nil
}

func styleAreaColumn(file *excelize.File) error {
	/*
		Set style into "Area" column
	*/

	// Define cell style
	styleArea, err := file.NewStyle(&excelize.Style{
		NumFmt:        4,
		DecimalPlaces: 2,
	})
	if err != nil {
		return err
	}

	// Apply modifications
	err = file.SetColStyle("Sheet1", "C", styleArea)
	if err != nil {
		return err
	}

	return nil
}
