package main

import (
	"encoding/base64"
	"log"
	"os/exec"
	"path/filepath"
)

func ExtractText(path string) string {
	fileName, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}

	f := "-f=" + "\"" + fileName + "\""
	l := "-l=" + "\"" + "eng" + "\""
	cmd := exec.Command(Config.Exec, Config.ExtractorPath, f, l)
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
