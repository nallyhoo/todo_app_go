package handlers

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "time"
    "todo-app/db"
    "todo-app/models"
    "github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, title, description, completed, created_at, start_time, end_time, progress FROM todos ORDER BY created_at DESC")
    if err != nil {
        log.Printf("Error querying todos: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var t models.Todo
        var startTime, endTime sql.NullTime
        err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &startTime, &endTime, &t.Progress)
        if err != nil {
            log.Printf("Error scanning todo: %v", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.StartTime = startTime.Time
        t.EndTime = endTime.Time
        todos = append(todos, t)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Printf("Invalid todo ID: %v", err)
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    var t models.Todo
    var startTime, endTime sql.NullTime
    err = db.DB.QueryRow("SELECT id, title, description, completed, created_at, start_time, end_time, progress FROM todos WHERE id = ?", id).
        Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &startTime, &endTime, &t.Progress)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            log.Printf("Error querying todo: %v", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    t.StartTime = startTime.Time
    t.EndTime = endTime.Time

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var t models.Todo
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        log.Printf("Error decoding request: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if t.Title == "" {
        http.Error(w, "Title is required", http.StatusBadRequest)
        return
    }

    if t.Progress < 0 || t.Progress > 100 {
        http.Error(w, "Progress must be between 0 and 100", http.StatusBadRequest)
        return
    }

    var startTime, endTime interface{}
    if !t.StartTime.IsZero() {
        startTime = t.StartTime
    } else {
        startTime = nil
    }
    if !t.EndTime.IsZero() {
        endTime = t.EndTime
    } else {
        endTime = nil
    }

    res, err := db.DB.Exec("INSERT INTO todos (title, description, completed, start_time, end_time, progress) VALUES (?, ?, ?, ?, ?, ?)",
        t.Title, t.Description, t.Completed, startTime, endTime, t.Progress)
    if err != nil {
        log.Printf("Error inserting todo: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := res.LastInsertId()
    t.ID = int(id)
    t.CreatedAt = time.Now()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(t)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Printf("Invalid todo ID: %v", err)
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    var t models.Todo
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        log.Printf("Error decoding request: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if t.Progress < 0 || t.Progress > 100 {
        http.Error(w, "Progress must be between 0 and 100", http.StatusBadRequest)
        return
    }

    var startTime, endTime interface{}
    if !t.StartTime.IsZero() {
        startTime = t.StartTime
    } else {
        startTime = nil
    }
    if !t.EndTime.IsZero() {
        endTime = t.EndTime
    } else {
        endTime = nil
    }

    _, err = db.DB.Exec("UPDATE todos SET title = ?, description = ?, completed = ?, start_time = ?, end_time = ?, progress = ? WHERE id = ?",
        t.Title, t.Description, t.Completed, startTime, endTime, t.Progress, id)
    if err != nil {
        log.Printf("Error updating todo: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.ID = id
    t.CreatedAt = time.Now()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Printf("Invalid todo ID: %v", err)
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    _, err = db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        log.Printf("Error deleting todo: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}