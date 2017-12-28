package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config reads the config
type Config struct {
	IP      string      `json:"ip"`
	Port    string      `json:"port"`
	Mongodb MongoConfig `json:"mongodb"`
}
type MongoConfig struct {
	IP       string `json:"ip"`
	Database string `json:"database"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
