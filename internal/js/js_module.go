package js

import (
	"ai-coding-comparison/internal"
	"embed"
	"net/http"
)

//go:embed js/*.js
var files embed.FS

func Init(app internal.App) {
	fs := http.FileServer(http.FS(files))
	app.Mux.Handle("/js/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		fs.ServeHTTP(w, r)
	}))
}
