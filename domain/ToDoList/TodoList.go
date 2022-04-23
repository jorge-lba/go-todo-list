package ToDoList

import (
	"time"

	"go-todo-list/domain/ToDo"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Props struct {
	gorm.Model
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Items       map[string]*ToDo.ToDo `json:"items,omitempty"`
	CreatedAt   time.Time             `json:"createdAt"`
	UpdatedAt   time.Time             `json:"updatedAt"`
	DeletedAt   time.Time             `json:"deletedAt,omitempty"`
}

type Entity struct {
	id    string
	props Props
}

// Create é semelhante há um método abstrato
func Create(p Props) Entity {
	t := Entity{}
	t.props = p

	if len(p.ID) > 0 {
		t.id = p.ID
	} else {
		t.id = uuid.New().String()
	}

	if p.CreatedAt.IsZero() {
		t.props.CreatedAt = time.Now().UTC()
	}

	if p.UpdatedAt.IsZero() {
		t.props.UpdatedAt = time.Now().UTC()
	}

	return t
}

// GetProps é semelhante há um método publico
func (t Entity) Id() string {
	return t.id
}

func (t Entity) Props() Props {
	return t.props
}

func (tl *Entity) IncludeItem(i ToDo.ToDo) string {
	t := ToDo.CreateToDo(i)
	tdId := t.Id
	tl.props.Items[tdId] = t

	return tdId
}

func (t *Entity) GetItems() map[string]*ToDo.ToDo {
	return t.props.Items
}

func (t *Entity) DoneToDo(i string) {
	t.props.Items[i].SetDone()
}

func (t *Entity) UndoneToDo(i string) {
	t.props.Items[i].SetUndone()
}
