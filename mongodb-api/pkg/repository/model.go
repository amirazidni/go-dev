package repository

type Todo struct {
	ID   string `bson:"_id" json:"id"`
	Name string `json:"name"`
}

type NewTodo struct {
	Name string `json:"name"`
}
