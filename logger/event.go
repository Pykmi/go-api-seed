package logger

import "fmt"

// Event : The Event struct
type Event struct {
	RemoteAddr string
	Status     int
	Method     string
	RequestURI string
}

// Print : Print event
func (e *Event) Print() {
	fmt.Printf("%v %v %v %v\n", e.Method, e.RemoteAddr, e.Status, e.RequestURI)
}

// NewEvent : Creates a new event
func NewEvent() *Event {
	e := &Event{}
	return e
}
