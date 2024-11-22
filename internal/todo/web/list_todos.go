package web

import (
	"ai-coding-comparison/internal"
	"ai-coding-comparison/internal/todo/db"
	"ai-coding-comparison/internal/todo/domain"
	"encoding/json"
	"net/http"
)

func ListTodos(app internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		todos, err := db.AllTodos(r.Context(), app.DB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonData := []domain.Todo{}
		for _, todo := range todos {
			jsonData = append(jsonData, domain.Todo{Title: todo.Title})
		}
		json.NewEncoder(w).Encode(jsonData)
	}
}
