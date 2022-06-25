package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const path = "/Users/roinuj16/Documents/JRempixels/02-VIAGENS/messy_files/2022-06"

func main() {

	if path != "" {

		files, err := ioutil.ReadDir(path)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		for _, file := range files {
			if !file.IsDir() {
				extension := filepath.Ext(file.Name())
				extension = strings.Trim(extension, ".")

				savingFolder := path + "/" + extension

				if _, err := os.Stat(savingFolder); os.IsNotExist(err) {
					os.Mkdir(savingFolder, os.ModePerm)
				}

				filename := path + "/" + file.Name()
				newFilename := savingFolder + "/" + file.Name()
				err := os.Rename(filename, newFilename)

				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			}
		}
	} else {
		fmt.Println("filesPath is empty")
	}
}
