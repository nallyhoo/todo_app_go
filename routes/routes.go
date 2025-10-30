package routes

import (
    "net/http"
    "todo-app/handlers"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
    r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

    // CORS middleware (simple for dev)
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            next.ServeHTTP(w, r)
        })
    })

    return r
}