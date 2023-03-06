package main

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	/* outputPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} */
	fileName, err := filepath.Abs("1.png")
	if err != nil {
		fmt.Println(err)
	}

	f := "-f=" + "\"" + fileName + "\""
	//o := "-o=" + "\"" + outputPath + "\""

	pathToExtractor := "../Python/Extractor/output/main/main.exe"
	fmt.Println(pathToExtractor)
	cmd := exec.Command("powershell", pathToExtractor, f)

	out, err := cmd.Output()
	//fmt.Println(string(out))
	if err != nil {
		fmt.Println(err)
	}

	data, err := base64.StdEncoding.DecodeString(string(out))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", data)
}
