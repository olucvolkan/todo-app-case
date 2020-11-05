package schema

type TodoListSchema struct {
	Todo       Todo       `json:"todo"`
	Inprogress InProgress `json:"in-progress"`
	Done       Done       `json:"done"`
}
type Todo struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type InProgress struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type Done struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type Items struct {
	Item []string
}

type Item struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
