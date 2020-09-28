package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler(c <-chan os.Signal) {
	s := <-c
	fmt.Println("Got signal: ", s)
}

func businessLogic() {
	fmt.Println(fmt.Sprintf("[%d]", os.Getpid()), "App is ready to receive signals!")
	for {}
}

func main() {
	c := make(chan os.Signal)
	go signalHandler(c)

	signal.Notify(c, syscall.SIGUSR1)
	signal.Ignore(syscall.SIGINT)

	businessLogic()
}
