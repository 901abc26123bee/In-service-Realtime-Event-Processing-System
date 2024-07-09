package main

import (
	"context"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// init observer, eventDispatcher and worker manager
	observer := NewEventObserver()
	eventDispatcher := NewEventDispatcher()
	worker1Manager := NewWorker1Manager()
	worker2Manager := NewWorker2Manager()

	// register/binding
	eventDispatcher.RegisterEventChan(observer.ChanEvent1, []chan TaskEventMessage{worker1Manager.Trigger, worker2Manager.Trigger})
	eventDispatcher.RegisterEventChan(observer.ChanEvent2, []chan TaskEventMessage{worker1Manager.Trigger})
	eventDispatcher.RegisterEventChan(observer.ChanEvent3, []chan TaskEventMessage{worker2Manager.Trigger})

	// run observer
	wg.Add(1)
	go func() {
		defer wg.Done()
		observer.SetUpListeners(ctx)
	}()

	// run eventDispatcher
	wg.Add(1)
	go func() {
		defer wg.Done()
		eventDispatcher.DispatchEvent(ctx)
	}()

	// run worker1Manager
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker1Manager.ManageWorker1(ctx)
	}()

	// run worker2Manager
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker2Manager.ManageWorker2(ctx)
	}()

	// wait till all goroutine to complete
	wg.Wait()
}

