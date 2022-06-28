package todo

// used to store todos
var TodoItems = []Todo{}

// used to assign unique id to each todo
var count = 1

// Here comes all the Queries

// This method is used to get all todos
func GetTodos() []Todo {
	return TodoItems
}

// This method is used to get a todo by id
// returns empty todo if not found
func GetTodo(id int) Todo {
	for _, todo := range TodoItems {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}

// Here comes all the Mutations

// This method is used to add a todo
// takes title as input and returns the todo
// returns the todo with assigned id
func AddTodo(title string) *Todo {
	// Create a new todo
	temp := &Todo{
		ID:        count,
		Title:     title,
		Completed: false,
	}
	// Add it to the list of todos
	TodoItems = append(TodoItems, *temp)
	// Increment the count
	count++
	// Return the todo
	return temp
}

// This method is used to update a todo
// takes id, title and completed as input
// returns true if updated successfully
// returns false if not found
func UpdateTodo(id int, title string, completed bool) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems[i].Title = title
			TodoItems[i].Completed = completed
			return true
		}
	}
	return false
}

// This method is used to delete a todo
// takes id as input
// returns true if deleted successfully
// returns false if not found
func DeleteTodo(id int) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems = append(TodoItems[:i], TodoItems[i+1:]...)
			return true
		}
	}
	return false
}
