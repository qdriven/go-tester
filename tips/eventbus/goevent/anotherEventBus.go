package goevent

import "fmt"

// Event can take all arguments of an event
type Event interface{}

// EventType is event type for an event :3
type EventType int

// EventObject contains event and event type
type EventObject struct {
	EventType
	Event
}

// EventHandler is handler for events and takes any arguments
type EventHandler func(args Event)

// register event types, this is just a format, don't register here,
// rather use your own types like in the example/events.go
const (
	NoobEvent EventType = iota
)

var eventMap map[EventType][]EventHandler
func init() {
	eventMap = make(map[EventType][]EventHandler)
}

// Subscribe to a function
func Subscribe(eH EventHandler, eT EventType) {
	if len(eventMap[eT]) == 0 {
		eventMap[eT] = make([]EventHandler, 0)
	}
	eventMap[eT] = append(eventMap[eT], eH)
}

// Publish event
func Publish(eT EventType, e Event) {
	for _, eH := range eventMap[eT] {
		eH(e)
	}
}

// Run is a goroutine for receving and publishing events
func Run(publisher chan EventObject) {
	for {
		eventObject := <-publisher
		fmt.Println("Event received ", eventObject.EventType)
		Publish(eventObject.EventType, eventObject.Event)
	}
}

// NewEventPublisher makes a new publisher channel for events
func NewEventPublisher() chan EventObject {
	publisher := make(chan EventObject)
	go Run(publisher)
	return publisher
}