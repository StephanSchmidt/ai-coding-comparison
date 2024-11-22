package todo

import (
	"ai-coding-comparison/internal"
	"ai-coding-comparison/internal/todo/web"
)

func Init(app internal.App) {
	app.Mux.HandleFunc("/", web.Index())
	app.Mux.HandleFunc("GET /list/", web.ListTodos(app))
	app.Mux.HandleFunc("POST /add/", web.AddTodo(app))
}
