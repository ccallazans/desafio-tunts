package main

type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	
	Capital    []string               `json:"capital"`
	Area       float64                `json:"area"`
	Currencies map[string]interface{} `json:"currencies"`
}

type CountryArray []Country
