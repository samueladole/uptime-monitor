package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
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
		message := fmt.Sprintf("[UP] %s\n", url)
		color.Green(message)
	} else {
		logStatus(url, false, resp.Status)
		message := fmt.Sprintf("[DOWN] %s : %s\n", url, resp.Status)
		color.Red(message)
	}
}
