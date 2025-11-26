package main

import (
	"log"
	"net/http"
	"time"
)

func checkWebsite(url string) {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		logStatus(url, false, err.Error())
		log.Printf("[DOWN] %s : %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		logStatus(url, true, "")
		log.Printf("[UP] %s\n", url)
	} else {
		logStatus(url, false, resp.Status)
		log.Printf("[DOWN] %s : %s\n", url, resp.Status)
	}
}
