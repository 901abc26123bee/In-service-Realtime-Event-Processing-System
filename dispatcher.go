package main

import (
	"context"
	"sync"
	"time"
)

// EventDispatcher defines struct of resource manager
type EventDispatcher struct {
	BindingMap map[chan TaskEventMessage][]chan TaskEventMessage
}

// NewEventDispatcher init and return a EventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		BindingMap: make(map[chan TaskEventMessage][]chan TaskEventMessage),
	}
}

// WatcherEventMessage defines watcher event message
type TaskEventMessage struct {
	TaskID      string
	Event       EventType
	Status      EventStatusType
	Description string
	TriggerTime time.Time
}

func (d *EventDispatcher) RegisterEventChan(sourceChan chan TaskEventMessage, receivers []chan TaskEventMessage) {
	if _, ok := d.BindingMap[sourceChan]; !ok {
		d.BindingMap[sourceChan] = []chan TaskEventMessage{}
	}
	d.BindingMap[sourceChan] = append(d.BindingMap[sourceChan], receivers...)
}

// BroadcastChanEvent broadcast source channel event message to register receiver channel
func (d *EventDispatcher) BroadcastChanEvent(sourceChan chan TaskEventMessage, receivers []chan TaskEventMessage) {
	for message := range sourceChan {
		// Use a WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup

		// Broadcast the message to all receiver channels concurrently
		for _, receiver := range receivers {
			wg.Add(1)
			go func(ch chan TaskEventMessage, msg TaskEventMessage) {
				defer wg.Done()
				ch <- msg
			}(receiver, message)
		}

		// Wait for all goroutines to finish before processing the next message
		wg.Wait()
	}
}

// DispatchEvent dispatch event to register
func (d *EventDispatcher) DispatchEvent(ctx context.Context) {
	for source, receivers := range d.BindingMap {
		go d.BroadcastChanEvent(source, receivers)
	}
}
