package main

import (
	"context"
	"os"
	"os/signal"
	"queue-mgr/internal/datasum"
	"queue-mgr/internal/dispatcher"
	"queue-mgr/internal/jobs"
	"syscall"
)

func main() {

	// Graceful Shutdown
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()

	// Dispatching
	w := datasum.NewDataSumWorker()

	d := dispatcher.NewDispatcher(w, 10, 1000)
	d.Start(ctx)

	for i := 0; i < 100; i++ {
		job := &jobs.Job{}
		d.Add(job)
	}

	d.Wait()
}
