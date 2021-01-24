package goevent

import (
	"fmt"
	"testing"
)

// my event types, registering to govent
const (
	MyMessage EventType = iota
	MyWeather EventType = iota
)

// MessageEvent is an event type for messaging
type MessageEvent struct {
	Message string
}

// WeatherEvent is an event type for weather conditions
type WeatherEvent struct {
	Condition string
}


func TestAnotherEventBus(t *testing.T) {
	publisher := NewEventPublisher()

	Subscribe(ShowMessage, MyMessage)
	Subscribe(ShowWeather, MyWeather)

	publisher <- EventObject{
		EventType: MyMessage,
		Event:     MessageEvent{Message: "Hello World"},
	}
	publisher <- EventObject{
		EventType: MyWeather,
		Event:     WeatherEvent{Condition: "Sunny"},
	}
	publisher <- EventObject{
		EventType: MyWeather,
		Event:     MessageEvent{Message: "Hello World"},
	}
	publisher <- EventObject{
		EventType: MyMessage,
		Event:     WeatherEvent{Condition: "Rainy"},
	}

	select {}
}

// ShowMessage prints the message
func ShowMessage(e Event) {
	if e, ok := e.(MessageEvent); ok {
		fmt.Println(e.Message)
	}
}


// ShowWeather prints the weather condition
func ShowWeather(e Event) {
	if e, ok := e.(WeatherEvent); ok {
		fmt.Println(e.Condition)
	}
}