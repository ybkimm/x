package servers

import (
	"context"
	"net/http"
	"time"
)

const (
	httpReadTimeout     = 30 * time.Second
	httpWriteTimeout    = 2 * time.Minute
	httpShutdownTimeout = 2 * time.Minute
)

func HTTP(addr string, h http.Handler) error {
	s := &http.Server{
		Addr:    addr,
		Handler: h,

		ReadTimeout:  httpReadTimeout,
		WriteTimeout: httpWriteTimeout,
	}

	sigchan := notifyInterrupt()

	errchan := make(chan error, 1)
	go func() {
		err := s.ListenAndServe()
		if err == http.ErrServerClosed {
			err = nil
		}
		errchan <- err
	}()

	for {
		select {
		case <-sigchan:
			ctx, cancel := context.WithTimeout(
				context.Background(),
				httpShutdownTimeout,
			)
			s.Shutdown(ctx)
			cancel()

		case err := <-errchan:
			return err
		}
	}
}
