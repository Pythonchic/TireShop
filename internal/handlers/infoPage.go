package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Pythonchic/giga-project/internal/config"
	"github.com/Pythonchic/giga-project/internal/storage"
)

func NewInfoHandler(filename string, config config.Config) (handler http.HandlerFunc, err error) {
	var globalErr error
	products, err := storage.GetProducts(config)

	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		}, fmt.Errorf("error read %s file: %s, error: %w", config.StorageFileFormat, config.StorageFile, err)

	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(filename)
		if err !=  nil {
			globalErr = fmt.Errorf("error parse file %s: ", filename)
		}

		data := storage.FillData(products)
		tmpl.Execute(w, data)
	}, globalErr
}
