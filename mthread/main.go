package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
  _ "net/http/pprof"
)

func main() {
	x := make(chan int, 128)
	done := make(chan bool)
	var wg sync.WaitGroup
	n := 4032
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(15 * time.Millisecond)
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					x <- 22
				}
			}
		}()
	}
	for i:=0; i<n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case <-x:
				}
			}
		}()
	}
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	close(done)
	wg.Wait()
	fmt.Println("bye")
}
