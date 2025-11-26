package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	interval := flag.Int("interval", 60, "Check interval in seconds")
	urlsFile := flag.String("file", "", "Path to file with URLs")
	webMode := flag.Bool("web", false, "Run as web server")
	flag.Parse()

	urls := flag.Args()
	if *urlsFile != "" {
		fileUrls, err := loadURLsFromFile(*urlsFile)
		if err != nil {
			log.Fatal("Failed to read URLs file:", err)
		}
		urls = append(urls, fileUrls...)
	}

	if len(urls) == 0 && !*webMode {
		log.Fatal("No URLs provided. Use CLI args, --file, or --web.")
	}

	if *webMode {
		fmt.Println("Starting web server on :8080...")
		startWebServer() // see web.go
		return
	}

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	fmt.Println("Monitoring URLs:", urls)
	for {
		for _, url := range urls {
			go checkWebsite(url)
		}
		<-ticker.C
	}
}
