package storage

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
)

type Product struct {
	Name             string `json:"Name" xml:"Name"`
	Image            string `json:"Image" xml:"Image"`
	Cost             int    `json:"Cost" xml:"Cost"`
	Diameter         int    `json:"Diameter" xml:"Diameter"`
	Season           string `json:"Season" xml:"Season"`
	SeasonTranslated string `json:"SeasonTranslated" xml:"SeasonTranslated"`
	Width            int    `json:"Width" xml:"Width"`
	Profile          int    `json:"Profile" xml:"Profile"`
	Manufacturer     string `json:"Manufacturer" xml:"Manufacturer"`
}

// ReadProductsFromJSON читает продукты из JSON файла
func ReadProductsFromJSON(filename string) ([]Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// ReadProductsFromXML читает продукты из XML файла
func ReadProductsFromXML(filename string) ([]Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	type Catalog struct {
		Products []Product `xml:"Product"`
	}

	var catalog Catalog
	err = xml.Unmarshal(data, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog.Products, nil
}
