package main

import (
	"bufio"
	"os"
)

func loadURLsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			urls = append(urls, line)
		}
	}
	return urls, scanner.Err()
}
