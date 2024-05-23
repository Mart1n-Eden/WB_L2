package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Определение флагов командной строки
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	// Парсинг аргументов командной строки
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	// Установка контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), *timeoutFlag)
	defer cancel()

	// Установка соединения с сервером
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Запуск горутины для чтения данных из сокета и вывода их в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		cancel() // Отмена контекста при закрытии сокета
	}()

	// Запуск горутины для чтения данных из STDIN и записи их в сокет
	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close() // Закрытие сокета при завершении чтения STDIN
	}()

	// Ожидание сигнала о завершении программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigCh:
		fmt.Println("\nClosing connection...")
	case <-ctx.Done():
		fmt.Println("\nConnection timed out")
	}
}
