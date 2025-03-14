package fileio

import "os"
import "log"
import "strings"

func ReadFile(file string) string {
	text, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}

func ReadFiles() []string {
	files := []string{}
	fileNames, err := os.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileNames {
		if strings.Contains(file.Name(), ".") {
			files = append(files, file.Name())
		}
	}

	return files
}
