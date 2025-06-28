package storage

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/Pythonchic/tireshop/internal/config"
)

type Product struct {
	Name             string `json:"Name" xml:"Name"`
	Image            string `json:"Image" xml:"Image"`
	ImagePath        string
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

func FillData(products []Product, config config.Config) Data {
	if len(products) == 0 {
		return Data{}
	}

	data := Data{
		Products:      products,
		MinPrice:      uint16(products[0].Cost),
		MaxPrice:      uint16(products[0].Cost),
		Widths:        make([]int, 0),
		Profiles:      make([]int, 0),
		Diameters:     make([]int, 0),
		StoragePath:   config.Storage.Images,
		StorageImages: config.Storage.Images,
	}

	// Временные мапы для хранения уникальных значений
	widthsMap := make(map[int]bool)
	profilesMap := make(map[int]bool)
	diametersMap := make(map[int]bool)

	for _, product := range products {
		// Обновляем минимальную и максимальную цену
		if product.Cost < int(data.MinPrice) {
			data.MinPrice = uint16(product.Cost)
		}
		if product.Cost > int(data.MaxPrice) {
			data.MaxPrice = uint16(product.Cost)
		}

		// Собираем уникальные значения
		widthsMap[product.Width] = true
		profilesMap[product.Profile] = true
		diametersMap[product.Diameter] = true
	}

	// Преобразуем мапы в слайсы
	for width := range widthsMap {
		data.Widths = append(data.Widths, width)
	}
	for profile := range profilesMap {
		data.Profiles = append(data.Profiles, profile)
	}
	for diameter := range diametersMap {
		data.Diameters = append(data.Diameters, diameter)
	}

	// Сортируем слайсы для удобства
	sort.Ints(data.Widths)
	sort.Ints(data.Profiles)
	sort.Ints(data.Diameters)

	return data
}

func GetProducts(config config.Config) (Products []Product, err error) {
	var products []Product

	absPath := config.Storage.Path + "/" + config.Storage.File

	switch strings.ToLower(config.Storage.FileFormat) {
	case "json":
		products, err = ReadProductsFromJSON(absPath)
	case "xml":
		products, err = ReadProductsFromXML(absPath)
	default:
		products, err = ReadProductsFromJSON(absPath)
	}


	for i, v := range(products) {
		products[i].ImagePath = config.Storage.Path + "/" + config.Storage.Images + "/" + v.Image
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}
