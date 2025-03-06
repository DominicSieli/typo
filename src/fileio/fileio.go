package fileio

import "os"
import "log"
import "strings"

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func PopulateFileList() []string {
	var files []string
	fileNames, err := os.ReadDir("./")
	extensions := []string{".c",".h",".cpp",".asm",".cs",".zig",".txt",".go",".js",".ts",".json",".java",".lua",".md",".py",".html",".css",".sh"}

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileNames {
		for _, extension := range extensions {
			if strings.Contains(file.Name(), extension) {
				files = append(files, file.Name())
			}
		}
	}

	return files
}
