package ToDo

import "github.com/google/uuid"

type ToDo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func CreateToDo(t ToDo) *ToDo {
	return &ToDo{
		Id:          uuid.New().String(),
		Title:       t.Title,
		Description: t.Description,
		Done:        false,
	}
}

func (t *ToDo) SetDone() {
	t.Done = true
}

func (t *ToDo) SetUndone() {
	t.Done = false
}
