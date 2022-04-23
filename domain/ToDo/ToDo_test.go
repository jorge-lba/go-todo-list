package ToDo_test

import (
	"go-todo-list/domain/ToDo"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// https://onsi.github.io/ginkgo/

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "To-Do Suite")
}

var _ = Describe("To-Do", func() {
	It("should create a new to-do", func() {
		toDo := ToDo.CreateToDo(ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
		})

		Expect(toDo.Title).To(Equal("First Test"))
		Expect(toDo.Description).To(Equal("Create a new To-Do"))
		Expect(toDo.Done).To(BeFalse())
		Expect(reflect.TypeOf(toDo.Id).Name()).To(Equal("string"))
	})

	It("should done to-do", func() {
		toDo := ToDo.CreateToDo(ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
		})

		toDo.SetDone()

		Expect(toDo.Done).To(BeTrue())
	})

	It("should be undone to-do", func() {
		toDo := ToDo.CreateToDo(ToDo.ToDo{
			Title:       "First Test",
			Description: "Create a new To-Do",
			Done:        true,
		})

		toDo.SetUndone()

		Expect(toDo.Done).To(BeFalse())
	})
})
