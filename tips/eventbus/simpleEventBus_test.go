package eventbus

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var eb = &SimpleEventBus{
	subscribers: map[string]DataChannelSlice{},
}

func publishTo(topic string, data interface{}) {
	eb.Publish(topic, data)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

type DataSample struct {
	Sample string
	Name   string
}

type Subscriber interface {
	onData(event DataEvent)
}

type MySubscriber struct {
}

func (m MySubscriber) onData(event DataEvent) {
	// do anything with event
	fmt.Printf("Topic: %s; DataEvent: %v\n", event.Topic, event.Data)
}

func TestPublish(t *testing.T) {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic3", ch3)

	go publishTo("topic1", "Hi topic 1")
	go publishTo("topic2", "Welcome to topic 2")
	go publishTo("topic3", &DataSample{Sample: "Sample",Name:"Name"})
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
