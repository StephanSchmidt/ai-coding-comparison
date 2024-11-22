package db

import (
	"context"

	"gorm.io/gorm"
)

func CreateTodo(ctx context.Context, db *gorm.DB, title string) error { 
	todo := TodoDB{Title: title}
	res := db.WithContext(ctx).Create(&todo)
	return res.Error
}

func AllTodos(ctx context.Context, db *gorm.DB) ([]TodoDB, error) {
	var todos []TodoDB
	res := db.WithContext(ctx).Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}
	return todos, nil
}
  
