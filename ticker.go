package main

import (
	"fmt"
	"time"
)

func startTicker(f func()) chan bool {
	done := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()
	return done
}

func mainTicker() {
	tickerGorutine := startTicker(func() {
		fmt.Println("tick...")
	})
	time.Sleep(5 * time.Second)
	close(tickerGorutine)
	time.Sleep(5 * time.Second)
}
