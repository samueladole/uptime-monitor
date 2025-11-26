package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func logStatus(url string, up bool, info string) {
	db, err := sql.Open("sqlite3", "./db/uptime.db")
	if err != nil {
		log.Println("DB error:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS status_logs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            url TEXT,
            up BOOLEAN,
            info TEXT,
            checked_at DATETIME
        )
    `)
	if err != nil {
		log.Println("Failed to create table:", err)
		return
	}

	_, err = db.Exec(`
        INSERT INTO status_logs(url, up, info, checked_at)
        VALUES (?, ?, ?, ?)`,
		url, up, info, time.Now())
	if err != nil {
		log.Println("Insert error:", err)
	}
}
