package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IPampurin/WebSocket-Server/pkg/server"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go signalListener(ctx, cancel)

	if err := server.Run(ctx); err != nil {
		fmt.Printf("Сервер завершился с ошибкой: %v", err)
		return
	}

	fmt.Println("Приложение крректно остановлено.")
}

// signalListener обрабатывает сигналы отмены
func signalListener(ctx context.Context, cancel context.CancelFunc) {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case <-ctx.Done():
		return
	case <-sigCh:
		cancel()
		return
	}
}
