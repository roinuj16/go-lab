package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	/**
	 *To use the file organizer just change the filesPath constant to point to the directory where the files are
	 * e.g /Users/user-name/folder
	 */
	const filesPath = "./assets"

	if filesPath != "" {
		organizeFilesByDay(filesPath)

	} else {
		fmt.Println("filesPath is empty")
	}
}

func organizeFilesByDay(path string) {
	fmt.Println("Start organizing by date...")
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		filetime := file.ModTime()
		savingFolder := path + "/" + filetime.Format("2006-01-02")

		fmt.Println("Date:", savingFolder)

		if _, err := os.Stat(savingFolder); os.IsNotExist(err) {
			os.Mkdir(savingFolder, os.ModePerm)
		}

		filename := path + "/" + file.Name()
		newFileName := savingFolder + "/" + file.Name()
		err := os.Rename(filename, newFileName)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		organizeFilesByExtension(savingFolder)

	}
}

func organizeFilesByExtension(folder string) {
	fmt.Println("Start organizing by extension...")
	files, err := ioutil.ReadDir(folder)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() {
			extension := filepath.Ext(file.Name())
			extension = strings.Trim(extension, ".")

			fmt.Println("Extension:", extension)

			savingFolder := folder + "/" + extension

			if _, err := os.Stat(savingFolder); os.IsNotExist(err) {
				os.Mkdir(savingFolder, os.ModePerm)
			}

			filename := folder + "/" + file.Name()
			newFilename := savingFolder + "/" + file.Name()
			err := os.Rename(filename, newFilename)

			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		}
	}
}
