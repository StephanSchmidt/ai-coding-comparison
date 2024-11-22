package db

type TodoDB struct {
	Id    int
	Title string
}

func (t *TodoDB) TableName() string {
	return "todos"
}
