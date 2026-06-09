package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	portServer = "8081"
	hostServer = "localhost"
)

func Run(ctx context.Context) error {

	// Инициализируем api

	// Используем

	// Создаём и настраиваем сервер
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", hostServer, portServer),
	}

	idleConnsClosed := make(chan struct{})

	go func() {

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		<-sigChan

		if err := srv.Shutdown(context.Background()); err != http.ErrServerClosed {

		}
	}()

	<-idleConnsClosed

	return nil
}
