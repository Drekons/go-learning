package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type server struct {
	status bool
	sync.Mutex
}

func main() {
	s := server{status: true}

	go squaring(&s)
	eventShutdown(&s)
}

func eventShutdown(s *server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	s.Lock()
	s.status = false
	s.Unlock()
	fmt.Println("Выхожу из программы")
}

func squaring(s *server) {
	n := 0
	for {
		if !s.status {
			break
		}
		s.Lock()
		n++
		if n%2 != 0 {
			s.Unlock()
			continue
		}
		fmt.Printf("Квадрат числа %d = %d\n", n, n*n)
		s.Unlock()
	}
}
