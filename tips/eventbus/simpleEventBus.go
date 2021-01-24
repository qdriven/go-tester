package eventbus

import "sync"

//pub/sub pattern

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent
type DataChannelSlice []DataChannel

// EventBus stores the information about subscribers interested for // a particular topic
type SimpleEventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

/**
Subscribe a topic
*/
func (eb *SimpleEventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		//INIT a subscribers
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *SimpleEventBus) Publish(topic string, data interface{}) {
	eb.rm.Lock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, slice DataChannelSlice) {
			for _, channel := range slice {
				channel <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.Unlock()
}
