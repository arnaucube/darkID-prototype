package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config reads the config
type Config struct {
	Port   string `json:"port"`
	Server string `json:"server"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
