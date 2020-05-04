package app

type Todo struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

type Input struct {
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

type Update struct {
	Message  *string `json:"message"`
	Complete *bool   `json:"complete"`
}

// Update update todo if data is not nil
// It can partially update or update all data
func (t *Todo) Update(update Update) {
	if update.Message != nil {
		t.Message = *update.Message
	}

	if update.Complete != nil {
		t.Complete = *update.Complete

	}
}

// convert Input struct to Todo struct
func (i *Input) toTodo() Todo {
	return Todo{
		Message:  i.Message,
		Complete: i.Complete,
	}
}
