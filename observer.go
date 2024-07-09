package main

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// EventStatusType defines event status type
type EventStatusType string

const (
	EventStatusTypeSuccess = "SUCCESS"
	EventStatusTypeFailed  = "FAILED"
)

// EventType defines event type
type EventType string

const (
	EventTypeUserRegister   = "USER_REGISTER"
	EventTypeUserUnRegister = "USER_UNREGISTER"
)

// EventObserver defines struct for event obsever
type EventObserver struct {
	ChanEvent1 chan TaskEventMessage
	ChanEvent2 chan TaskEventMessage
	ChanEvent3 chan TaskEventMessage
}

// NewEventObserver init and return a new EventObserver 
func NewEventObserver() *EventObserver {
	return &EventObserver{
		ChanEvent1: make(chan TaskEventMessage, 50),
		ChanEvent2: make(chan TaskEventMessage, 50),
		ChanEvent3: make(chan TaskEventMessage, 10),
	}
}

// SetUpListeners set up event listener
func (o *EventObserver) SetUpListeners(ctx context.Context) {
	go o.listenToEvent1()
	go o.listenToEvent2()
	go o.listenToEvent3()
}

func (o *EventObserver) listenToEvent1() {
	mockEvent := time.NewTicker(2 * time.Second)
	for t := range mockEvent.C {
		o.ChanEvent1 <- TaskEventMessage{TaskID: uuid.New().String(), Event: EventTypeUserRegister, Description: "event1", TriggerTime: t}
	}
}

func (o *EventObserver) listenToEvent2() {
	mockEvent := time.NewTicker(2 * time.Second)
	for t := range mockEvent.C {
		o.ChanEvent2 <- TaskEventMessage{TaskID: uuid.New().String(), Event: EventTypeUserRegister, Description: "event2", TriggerTime: t}
	}
}

func (o *EventObserver) listenToEvent3() {
	mockEvent := time.NewTicker(2 * time.Second)
	for t := range mockEvent.C {
		o.ChanEvent3 <- TaskEventMessage{TaskID: uuid.New().String(), Event: EventTypeUserUnRegister, Description: "event3", TriggerTime: t}
	}
}
