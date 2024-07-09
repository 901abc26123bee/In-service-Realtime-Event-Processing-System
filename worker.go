package main

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Worker1Manager defines struct for Worker1Manager
type Worker1Manager struct {
	running bool
	mutex   sync.Mutex
	Trigger chan TaskEventMessage
}
// Worker2Manager defines struct for Worker2Manager
type Worker2Manager struct {
	running bool
	mutex   sync.Mutex
	Trigger chan TaskEventMessage
}

// NewWorker1Manager init and return a new Worker1Manager 
func NewWorker1Manager() *Worker1Manager {
	return &Worker1Manager{
		Trigger: make(chan TaskEventMessage, 20),
	}
}

// NewWorker2Manager init and return a new Worker2Manager 
func NewWorker2Manager() *Worker2Manager {
	return &Worker2Manager{
		Trigger: make(chan TaskEventMessage, 20),
	}
}

// ManageWorker1 manage worker1 process when receive event from trigger channel
func (m *Worker1Manager) ManageWorker1(ctx context.Context) {
	for message := range m.Trigger {
		go m.worker1Process(context.Background(), &message)
	}
}

// ManageWorker2 manage worker2 process when receive event from trigger channel
func (m *Worker2Manager) ManageWorker2(ctx context.Context) {
	for message := range m.Trigger {
		go m.worker2Process(context.Background(), &message)
	}
}

func (m *Worker1Manager) worker1Process(ctx context.Context, event *TaskEventMessage) {
	if m.running {
		return // skip turns if process already in running state
	}

	m.mutex.Lock()
	defer func() {
		m.running = false
		m.mutex.Unlock()
	}()

	if m.running {
		return // skip turns if process is running
	}
	m.running = true

	execTime := time.Now()
	time.Sleep(1 * time.Second) // processing
	log.Infof("successfully finish worker 1 process taskID %s, event %s description %s at with duration %s", event.TaskID, event.Event, event.Description , time.Since(execTime))
	return
}

func (m *Worker2Manager) worker2Process(ctx context.Context, event *TaskEventMessage) {
	if m.running {
		return // skip turns if process already in running state
	}

	m.mutex.Lock()
	defer func() {
		m.running = false
		m.mutex.Unlock()
	}()

	if m.running {
		return // skip turns if process is running
	}
	m.running = true

	execTime := time.Now()
	time.Sleep(2 * time.Second) // processing
	log.Infof("successfully finish worker 2 process taskID %s, event %s description %s at with duration %s", event.TaskID, event.Event, event.Description , time.Since(execTime))
	return
}
