package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Printf("Got %v...\n", <-interrupt)

	// здесь выключаем веб-сервер
	// и любые другие обработчики задач

}
