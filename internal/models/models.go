package models

// Struct to parse API data
type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	Capital    []string               `json:"capital"`
	Area       float64                `json:"area"`
	Currencies map[string]interface{} `json:"currencies"`
}

// Struct to parse row data
type HeaderRow struct {
	Sheet string
	Cell  string
	Text  string
}
