package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID string `json:"id"`
}

type Task struct {
	ID          string      `json:"id"`          // The unique id of the task
	Title       string      `json:"title"`       // Title of the task
	Description string      `json:"description"` // The descroption of the task
	Assignee    []User      `json:"assignee"`    // The list of user assigned to that task
	SharedWith  []User      `json:"sharedWith"`  // Users that the task shared with
	Dates       []time.Time `json:"dates"`       // The array of dates, it contains all history of dates, the last one will be the final date
	Plannings   []string    `json:"plannings"`   // Planning IDs that contains this task
	Owners      []User      `json:"owners"`      // The owner of the task
	Done        bool        `json:"done"`        // A boolean indicating that wether a task is done or not
}

// NewTask creates new task with given paramters
func NewTask(title, description string, assignee, sharedWith []User, plannings []string, owners []User) Task {
	return Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Assignee:    assignee,
		SharedWith:  sharedWith,
		Plannings:   plannings,
		Owners:      owners,
		Done:        false,
	}
}

func (t *Task) ChangeTitle(newTitle string) {
	t.Title = newTitle
}

func (t *Task) ChangeDescription(description string) {
	t.Description = description
}

func (t *Task) AssignTo(user User) error {
	if t.hasAssignee(user) {
		return fmt.Errorf("Task %s is already assigned to user %s", t.ID, user.ID)
	}
	t.Assignee = append(t.Assignee, user)
	return nil
}

func (t *Task) UnassignFrom(assignee User) error {
	index, err := t.indexOfAssignee(assignee)
	if err != nil {
		return fmt.Errorf("Task %s is not assigned to user %s", t.ID, assignee.ID)
	}
	t.Assignee = removElementFromSliceeWithIndex(t.Assignee, index)
	return nil
}

func (t *Task) AddToPlanning(planningId string) error {
	if t.isInPlanning(planningId) {
		return fmt.Errorf("Task %s is already in planning %s", t.ID, planningId)
	}
	t.Plannings = append(t.Plannings, planningId)
	return nil
}

func (t *Task) RemoveFromPlanning(pID string) error {
	index := indexOf(t.Plannings, pID, func(s1, s2 string) bool { return s1 == s2 })
	if index == -1 {
		return fmt.Errorf("Task %s is not in the planning %s", t.ID, pID)
	}
	t.Plannings = removElementFromSliceeWithIndex(t.Plannings, index)
	return nil
}

// TODO: Add owner
// TODO: Remove owner
// TODO: Put the task in done
// TODO: Put the task to undone

func (t Task) hasAssignee(user User) bool {
	return isExistsInSlice(t.Assignee, user, func(e1 User, e2 User) bool { return e1.ID == e2.ID })
}

func (t Task) isInPlanning(planningId string) bool {
	return isExistsInSlice(t.Plannings, planningId, func(e1, e2 string) bool { return e1 == e2 })
}

func (t Task) indexOfAssignee(user User) (int, error) {
	index := indexOf(t.Assignee, user, func(u1, u2 User) bool { return u1.ID == u2.ID })
	if index != -1 {
		return index, nil
	}
	return -1, fmt.Errorf("Task %s is not assigned to user %s", t.ID, user.ID)
}
