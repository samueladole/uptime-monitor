package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func startWebServer() {
	http.HandleFunc("/", dashboardHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./db/uptime.db")
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT url, up, info, checked_at FROM status_logs ORDER BY checked_at DESC LIMIT 50`)
	if err != nil {
		http.Error(w, "Query error", 500)
		return
	}
	defer rows.Close()

	fmt.Fprintln(w, "<html><body><h1>Uptime Monitor</h1><table border='1'><tr><th>URL</th><th>Status</th><th>Info</th><th>Checked At</th></tr>")
	for rows.Next() {
		var url, info string
		var up bool
		var checkedAt string
		rows.Scan(&url, &up, &info, &checkedAt)
		status := "DOWN"
		if up {
			status = "UP"
		}
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", url, status, info, checkedAt)
	}
	fmt.Fprintln(w, "</table></body></html>")
}
