package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigST struct {
	Host           string `json:"host"`
	TranslatorPath string `json:"translator_path"`
	ExtractorPath  string `json:"extractor_path"`
	Exec           string `json:"exec_name"`
}

var Config = loadConfig()

func loadConfig() *ConfigST {
	var cfg ConfigST
	data, err := os.ReadFile("config.json")
	if err == nil {
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatalln("Config", err)
		}
	} else {
		host := ":8070"
		trlPath := "../Python/Extractor/output/extractor/extractor.exe"
		etrPath := "../Python/Translator/output/translator/translator.exe"
		exe := "powershell"
		cfg.Host = host
		cfg.TranslatorPath = trlPath
		cfg.ExtractorPath = etrPath
		cfg.Exec = exe
	}
	return &cfg
}
