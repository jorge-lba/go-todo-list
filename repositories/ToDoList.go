package iRepo

import toDoList "go-todo-list/domain/ToDoList"

type ToDoList interface {
	Save(tdl *toDoList.Entity)
}
