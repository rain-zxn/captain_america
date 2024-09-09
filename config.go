package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config *Config

type Config struct {
	CaptainWebApp string `json:"captain_web_app"`
	CaptainX      string `json:"captain_x"`
	CaptainGroup  string `json:"captain_group"`
	CaptainAbout  string `json:"captain_about"`
	TextWebApp    string `json:"text_web_app"`
	TextGroup     string `json:"text_group"`
	TextX         string `json:"text_x"`
	TextAbout     string `json:"text_about"`
}

func ReadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	config = new(Config)
	if err := json.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}
