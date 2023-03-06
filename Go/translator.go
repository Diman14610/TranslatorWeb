package main

import (
	"encoding/base64"
	"log"
	"os/exec"
)

func Translate(text string) string {
	t := "-t=" + "\"" + text + "\""
	pathToExtractor := "../Python/Translator/output/main/main.exe"
	cmd := exec.Command("powershell", pathToExtractor, t)

	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}

	data, err := base64.StdEncoding.DecodeString(string(out))
	if err != nil {
		log.Println(err)
	}

	return string(data)
}
