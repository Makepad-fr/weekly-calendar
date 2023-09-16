package core

import "time"

type User struct {
}

type Task struct {
	ID          string      `json:"id"`          // The unique id of the event
	Title       string      `json:"title"`       // Title of the event
	Description string      `json:"description"` // The descroption of the event
	Assignee    []User      `json:"assignee"`    // The list of user assigned to that event
	SharedWith  []User      `json:"sharedWith"`  // Users that the event shared with
	Dates       []time.Time `json:"dates"`       // The array of dates, it contains all history of dates, the last one will be the final date
	Calendars   []string    `json:"calendars"`   // Calendar IDs that contains this event
	Owners      []User      `json:"owners"`      // The owner of the event
	Done        bool        `json:"done"`        // A boolean indicating that wether a task is done or not
}
