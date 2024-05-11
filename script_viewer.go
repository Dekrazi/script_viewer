package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"
)

// cspell: disable

// CHANGE THE DIRECTORY OF YOUR SCRIPTS FOLDER
var scriptsFolder string = "/home/username/my_scripts/"

// cspell: enable

func readLastLineAsDesc(file *os.File) (string, error) {
	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return lastLine, nil
}

func listFiles(path string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				log.Fatalf("Error opening file: %v", err)
			}
			defer file.Close()

			lastLine, err := readLastLineAsDesc(file)
			if err != nil {
				log.Fatalf("Error reading last line of file %s: %v", info.Name(), err)
			}
			fmt.Fprintf(w, "%s\t%s\n", info.Name(), lastLine)
		}
		return nil
	})
}

func main() {
	listFiles(scriptsFolder)
}
