package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config Model
type Config struct {
	JWT struct {
		SigningKey string `json:"signing_key"`
	} `json:"jwt"`
	MySQL struct {
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
}

const configPath = "./config.json"

// CFG public config
var CFG Config

// Load config file
func Load() {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("loading config file ", err)
	}

	if err := json.Unmarshal(file, &CFG); err != nil {
		log.Fatal("loading config file ", err)
	}
}
