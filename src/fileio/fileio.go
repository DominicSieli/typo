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
	extensions := []string{".c",".h",".cpp",".asm",".cs",".zig",".txt",".go",".js",".ts",".json",".java",".lua",".md",".py",".html",".css",".sh"}

	if err != nil {
		log.Fatal(err)
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
