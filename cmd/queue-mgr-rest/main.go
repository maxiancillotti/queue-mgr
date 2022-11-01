package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"queue-mgr/app"
	"queue-mgr/app/config"
	"queue-mgr/internal/handlers"
	"queue-mgr/internal/handlers/middlewares"
	"queue-mgr/internal/presenter"
	"queue-mgr/internal/service/datasum"
	"queue-mgr/internal/service/dispatcher"
	"queue-mgr/internal/service/queue"

	"github.com/gorilla/mux"
)

const (
	configFileDirName        = "queue-mgr-rest"
	enableStrictSlashRouting = true
)

var (
	configData = config.GetConfig(configFileDirName)

	q = queue.NewQueuer()
	w = datasum.NewDataSumWorker()

	d = dispatcher.NewDispatcherBuilder().
		SetTimeBetweenJobProcesses(10*time.Second).
		BuildDispatcher(w, q)

	p = presenter.NewJSONPresenter()

	c = handlers.NewJobsController(d, q, p)

	prMDW   = middlewares.NewPanicRecoverMiddleware(p)
	authMDW = middlewares.NewAuthController(p)

	httpRouter = mux.NewRouter().StrictSlash(enableStrictSlashRouting)
)

func main() {

	/******** Graceful Shutdown ********/
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()

	d.Start(ctx)

	setApiRoutes()
	setRoutesBase()

	app.StartHttpServer(&configData.HttpServer, CaselessMatcher(httpRouter))

}
