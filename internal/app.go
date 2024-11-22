package internal

import (
	"net/http"

	"gorm.io/gorm"
)

type App struct {
	DB  *gorm.DB
	Mux *http.ServeMux
}
