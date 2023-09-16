package core

type Planning struct {
	ID          string `json:"id"`          // The unique ID of the calendar
	Events      []Task `json:"events"`      // The list of event associated to that calendar
	SharedWith  []User `json:"sharedWith"`  // The list of users with who this event is shared
	Owner       []User `json:"owners"`      // The owner of the calendar
	Title       string `json:"title"`       // The title of the calendar
	Description string `json:"description"` // The description of the calendar
}
