package inMemory

import (
	"go-todo-list/domain/ToDoList"
)

type ToDoListInMemory struct {
	entities map[string]*ToDoList.Entity
}

func (tdr *ToDoListInMemory) Save(tl *ToDoList.Entity) {
	id := tl.Id()
	tdr.entities[id] = tl
}

func (tdr *ToDoListInMemory) FindById(id string) *ToDoList.Entity {
	return tdr.entities[id]
}

var ToDoListRepo = ToDoListInMemory{
	entities: make(map[string]*ToDoList.Entity),
}
