package main

import (
    "log"
    "net/http"
    "todo-app/db"
    "todo-app/routes"
)

func main() {
    db.InitDB()
    defer db.DB.Close()

    router := routes.SetupRouter()

    // Serve static files
    fs := http.FileServer(http.Dir("frontend"))
    router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

    log.Println("Server starting on :http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", router))
}