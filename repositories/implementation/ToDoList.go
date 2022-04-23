package repo

import (
	"fmt"
	"go-todo-list/domain/ToDo"
	"go-todo-list/domain/ToDoList"
	model "go-todo-list/storage/gorm"

	"gorm.io/gorm"
)

type ToDoListRepo struct {
	DB *gorm.DB
}

func (repo *ToDoListRepo) FindById(id string) ToDoList.Entity {
	var todoList model.ToDoList
	repo.DB.Model(&todoList).Preload("Items").First(&todoList, "id = ?", id)

	var items = make(map[string]*ToDo.ToDo)

	for _, value := range todoList.Items {
		items[value.ID] = mapperToDoEntityToDomain(value)
	}

	fmt.Println(items)

	tdl := ToDoList.Create(ToDoList.Props{
		ID:          todoList.ID,
		Title:       todoList.Title,
		Description: todoList.Description,
		Items:       items,
		CreatedAt:   todoList.CreatedAt,
		UpdatedAt:   todoList.UpdatedAt,
		DeletedAt:   todoList.DeletedAt.Time,
	})

	return tdl
}

func (repo *ToDoListRepo) Save(tdl *ToDoList.Entity) {
	p := tdl.Props()
	var items []model.ToDo

	for _, value := range p.Items {
		items = append(items, mapperToDoEntityToPercistense(value))
	}

	fmt.Println(items)

	repo.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&model.ToDoList{
		ID:          tdl.Id(),
		Title:       p.Title,
		Description: p.Description,
		Items:       items,
	})
}

func mapperToDoEntityToPercistense(td *ToDo.ToDo) model.ToDo {
	return model.ToDo{
		ID:          td.Id,
		Title:       td.Title,
		Description: td.Description,
		Done:        td.Done,
	}
}

func mapperToDoEntityToDomain(td model.ToDo) *ToDo.ToDo {
	return &ToDo.ToDo{
		Id:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Done:        td.Done,
	}
}
