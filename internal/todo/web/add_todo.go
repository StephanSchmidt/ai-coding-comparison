package web

import (
	"ai-coding-comparison/internal"
	"ai-coding-comparison/internal/todo/domain"
	"ai-coding-comparison/internal/todo/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func AddTodo(app internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo domain.Todo
		json.NewDecoder(r.Body).Decode(&todo)
		log.Default().Println("got todo", todo)
		err := usecase.AddTodoUC(r.Context(), app.DB, todo.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
