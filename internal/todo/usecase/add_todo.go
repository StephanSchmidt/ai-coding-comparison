package usecase

import (
	"ai-coding-comparison/internal/todo/db"
	"context"

	"gorm.io/gorm"
)

func AddTodoUC(ctx context.Context, gdb *gorm.DB, title string) error {
	return db.CreateTodo(ctx, gdb, title)
}
