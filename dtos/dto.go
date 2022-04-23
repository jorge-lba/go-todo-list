package dtos

import (
	"go-todo-list/domain/ToDo"
	"go-todo-list/domain/ToDoList"
	"time"
)

type ToDoDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func ToDoEntityToDTO(td *ToDo.ToDo) ToDoDTO {
	return ToDoDTO{
		ID:          td.Id,
		Title:       td.Title,
		Description: td.Description,
		Done:        td.Done,
	}
}

type ToDoListDTO struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Items       []*ToDoDTO `json:"items,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   time.Time  `json:"deletedAt,omitempty"`
}

func ToDoListEntityToDTO(tdl *ToDoList.Entity) ToDoListDTO {
	p := tdl.Props()

	var items []*ToDoDTO
	for _, value := range p.Items {
		items = append(items, &ToDoDTO{
			ID:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			Done:        value.Done,
		})
	}

	return ToDoListDTO{
		ID:          tdl.Id(),
		Title:       p.Title,
		Description: p.Description,
		Items:       items,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}
