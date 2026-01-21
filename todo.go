package lerangolang

// структура элемента таблицы todo list
type TodoList struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// стрктура элемента таблицы todo item
type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// структура элемента таблицы связи User <->List
type UserList struct {
	Id     int
	UserId int
	ListId int
}

// структура элемента таблицы связи List <->Item
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
