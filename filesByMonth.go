package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const folder = "/Users/roinuj16/Documents/JRempixels/02-VIAGENS/messy_files"

func main() {

	if folder != "" {

		files, err := ioutil.ReadDir(folder)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		for _, file := range files {
			filetime := file.ModTime()
			savingFolder := folder + "/" + filetime.Format("2006-01")

			if _, err := os.Stat(savingFolder); os.IsNotExist(err) {
				os.Mkdir(savingFolder, os.ModePerm)
			}

			filename := folder + "/" + file.Name()
			newFileName := savingFolder + "/" + file.Name()
			err := os.Rename(filename, newFileName)

			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

		}
	} else {
		fmt.Println("folder is empty")
	}
}
