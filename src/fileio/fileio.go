package fileio

import "os"
import "log"
import "strings"

func ReadFile(file string) string {
	text, error := os.ReadFile(file)

	if error != nil {
		log.Fatal(error)
	}

	return string(text)
}

func ReadFiles() []string {
	var files []string
	fileNames, error := os.ReadDir("./")
	extensions := []string{".c",".h",".cpp",".asm",".cs",".zig",".txt",".go",".js",".ts",".json",".java",".lua",".md",".py",".html",".css",".sh"}

	if error != nil {
		log.Fatal(error)
	}

	for _, file := range fileNames {
		for _, extension := range extensions {
			if strings.Contains(file.Name(), extension) {
				files = append(files, file.Name())
				break
			}
		}
	}

	return files
}
