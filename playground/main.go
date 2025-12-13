package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "./"

	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error stating file:", err)
		return
	}

	if stat.IsDir() {
		fmt.Printf("%s is a directory\n", path)

		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("Directory is empty")
			return
		}

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return
		}
		for _, file := range files {
			fmt.Println(filepath.Join(cwd, file.Name()))
		}
	} else {
		fmt.Printf("%s is a file\n", path)

		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return
		}
		fmt.Printf("%s\n", filepath.Join(cwd, file.Name()))
	}
}
