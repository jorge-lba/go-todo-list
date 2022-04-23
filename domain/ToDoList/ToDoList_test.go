package ToDoList_test

import (
	"go-todo-list/domain/ToDo"
	"go-todo-list/domain/ToDoList"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// https://onsi.github.io/ginkgo/

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "To-Do List Suite")
}

var _ = Describe("To-Do List", func() {
	It("should create a new to-do", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			Title:       "Test List",
			Description: "Create To-Do List",
		})

		props := toDoList.Props()

		Expect(props.Title).To(Equal("Test List"))
		Expect(props.Description).To(Equal("Create To-Do List"))
		Expect(props.CreatedAt).To(Not(BeZero()))
		Expect(props.UpdatedAt).To(Not(BeZero()))
		Expect(props.DeletedAt).To(BeZero())
	})

	It("should include one to-do in to-do list", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			Title:       "Test List",
			Description: "Create To-Do List",
		})

		toDoData := ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
		}

		toDoList.IncludeItem(toDoData)

		Expect(len(toDoList.GetItems())).To(Equal(1))
	})

	It("should done to-do in list", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			Title:       "Test List",
			Description: "Create To-Do List",
		})

		toDoData := ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
		}

		toDoId := toDoList.IncludeItem(toDoData)
		toDoList.DoneToDo(toDoId)

		Expect(toDoList.GetItem(toDoId).Done).To(Equal(true))
	})

	It("should undone to-do in list", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			Title:       "Test List",
			Description: "Create To-Do List",
		})

		toDoData := ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
		}

		toDoId := toDoList.IncludeItem(toDoData)
		toDoList.DoneToDo(toDoId)
		toDoList.UndoneToDo(toDoId)

		Expect(toDoList.GetItem(toDoId).Done).To(Equal(false))
	})

	It("should get id to-do list", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			Title:       "Test List",
			Description: "Create To-Do List",
		})

		Expect(toDoList.Id()).To(Not(BeEmpty()))
	})

	It("should load to-do", func() {
		toDoList := ToDoList.Create(ToDoList.Props{
			ID:          uuid.NewString(),
			Title:       "Test List",
			Description: "Create To-Do List",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			DeletedAt:   time.Now().UTC(),
		})

		props := toDoList.Props()

		Expect(props.Title).To(Equal("Test List"))
		Expect(props.Description).To(Equal("Create To-Do List"))
		Expect(props.CreatedAt).To(Not(BeZero()))
		Expect(props.UpdatedAt).To(Not(BeZero()))
		Expect(props.DeletedAt).To(Not(BeZero()))
	})
})
