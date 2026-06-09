package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	portServer = "8081"
	hostServer = "localhost"
)

func Run(ctx context.Context) error {

	// создаём и настраиваем сервер
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", hostServer, portServer),
	}

	// errCh канал для ошибок при работе сервера
	errCh := make(chan error, 1)

	// запускаем сервер
	go func() {

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}

	}()

	// ждём либо отмены контекста и завершения сервера, либо появления ошибки
	select {
	case <-ctx.Done():

		ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		if err := srv.Shutdown(ctxShutdown); err != nil {
			return fmt.Errorf("ошибка при graceful shutdown сервера: %w", err)
		}

		return nil

	case err := <-errCh:

		return err
	}
}
