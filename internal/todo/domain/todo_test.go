package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoCreation(t *testing.T) {
	todo := Todo{
		Title: "Test Todo",
	}

	// Using testify's assert package for better readability
	assert.Equal(t, "Test Todo", todo.Title, "Title should be 'Test Todo'")
}
