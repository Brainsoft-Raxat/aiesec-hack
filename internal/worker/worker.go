package worker

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	ticker *time.Ticker
	fn     func(context.Context) error
}

func NewHandler(interval time.Duration, fn func(context.Context) error) *Handler {
	return &Handler{
		ticker: time.NewTicker(interval),
		fn:     fn,
	}
}

func Start(ctx context.Context, handlers []*Handler) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for _, handler := range handlers {
		wg.Add(1)

		go func(h *Handler) {
			defer wg.Done()

			// Immediately run fn once after startup
			err := h.fn(ctx)
			if err != nil {
				logrus.Error("Error in worker:", err)
			}
			logrus.Info("Worker Success")

			for {
				select {
				case <-ctx.Done():
					return // Exit the worker when the context is canceled
				case <-h.ticker.C:
					err := h.fn(ctx)
					if err != nil {
						logrus.Error("Error in worker:", err)
					}
					logrus.Info("Worker Success")
				}
			}
		}(handler)
	}

	select {
	case sig := <-sigCh:
		logrus.Infof("Received signal %s, initiating graceful shutdown...", sig)
		cancel()
	case <-ctx.Done():
	}

	wg.Wait()

	logrus.Info("Worker gracefully shut down.")
}

func New(service *service.Service) []*Handler {
	return []*Handler{
		{
			ticker: time.NewTicker(30 * time.Minute),
			fn:     service.EventService.FetchAndCache,
		},
	}
}
