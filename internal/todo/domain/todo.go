package domain
type Todo struct {
    Title string `json:"title"` // The title of the todo item
	Done  bool   `json:"done"`  // Indicates if the todo item is completed
}
