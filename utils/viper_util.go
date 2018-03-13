package utils

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func LoadConfUtil(fileName string) *viper.Viper {
	viper := viper.New()
	viper.SetConfigType("toml")
	path := BuffetPath
	if len(path) <= 0 {
		log.Fatalln("buffet_path is not set")
	}

	filePath := filepath.Join(path, fileName)
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error: %v\n", err)
	}
	return viper
}
