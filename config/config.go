package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host string
}

var CONF Config

func LoadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&CONF)
	if err != nil {
		fmt.Println("error:", err)
	}
}
