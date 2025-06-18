package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"


)
type Config struct {
	Address string `yaml:"Address" yaml-default:"localhost:8080"`
	StoragePath     string `yaml:"StoragePath"`
	StorageFileFormat string `yaml:"StorageFileFormat"`
	StorageFile     string `yaml:"StorageFile"`
}



func ParseConfig(filePath string) (Config, error){
	var config Config
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return config, fmt.Errorf("config file does not exist %s", filePath)
	}

	if err := cleanenv.ReadConfig(filePath, &config); err != nil {
		return config, fmt.Errorf("cannot read config %s", err)
	}

	return config, nil
}
