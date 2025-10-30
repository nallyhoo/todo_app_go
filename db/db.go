package db

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    
)

var DB *sql.DB

func InitDB() {
    var err error
    // Update DSN with your MySQL details: user:pass@tcp(host:port)/dbname
    DSN := "root:MyDnDb6939$@tcp(127.0.0.1:3306)/todo_db?parseTime=true"
    DB, err = sql.Open("mysql", DSN)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Database connected successfully!")
}