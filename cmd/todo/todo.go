package main

import (
	"log"
	"net/http"

	"ai-coding-comparison/internal"
	"ai-coding-comparison/internal/js"
	"ai-coding-comparison/internal/todo"
)

func main() {
	mux := http.NewServeMux()

	db := initDb()

	app := internal.App{DB: db, Mux: mux}

	todo.Init(app)
	js.Init(app)

	log.Println("starting server on port http://localhost:8090")
	http.ListenAndServe("localhost:8090", mux)
}
