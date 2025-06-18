package storage

import (
	"sort"
	"strings"

	"github.com/Pythonchic/giga-project/internal/config"
)

type Data struct {
	Products  []Product
	MinPrice  uint16
	MaxPrice  uint16
	Widths    []int
	Profiles  []int
	Diameters []int
}


func FillData(products []Product) Data {
	if len(products) == 0 {
		return Data{}
	}

	data := Data{
		Products:  products,
		MinPrice:  uint16(products[0].Cost),
		MaxPrice:  uint16(products[0].Cost),
		Widths:    make([]int, 0),
		Profiles:  make([]int, 0),
		Diameters: make([]int, 0),
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


func GetProducts(config config.Config) (Products []Product,  err error) {
	var products []Product

	absPath := config.StoragePath + "/" + config.StorageFile

	switch strings.ToLower(config.StorageFileFormat) {
	case "json":
		products, err = ReadProductsFromJSON(absPath)
	case "xml":
		products, err = ReadProductsFromXML(absPath)
	default:
		products, err = ReadProductsFromJSON(absPath)
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}
