package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"queue-mgr/internal/jobs"
	"queue-mgr/internal/service/datasum"
	"queue-mgr/internal/service/dispatcher"
)

func main() {

	/******** Graceful Shutdown ********/
	//ctx, cancel := context.WithCancel(context.Background())
	_, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()

	/******** INITIALIZING ********/
	w := datasum.NewDataSumWorker()

	// GET FROM OS.ENV or VIPER LIBRARY
	//maxWorkers := 10 // Should be no more than CPU cores
	//buffers := 1000  // Depends on memory

	//d := dispatcher.NewDispatcher(w, maxWorkers, buffers)
	//d.Start(ctx)

	d := dispatcher.NewDispatcherBuilder().BuildDispatcher(w)

	/******** Dispatching / Queueing ********/

	// Controller layer - Handlers
	for i := 0; i < 100; i++ {
		job := &jobs.Job{}

		// Service layer
		d.Enqueue(job)
	}

	// No need because of http server
	d.Wait()
}
