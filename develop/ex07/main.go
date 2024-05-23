package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	done := make(chan interface{})

	go func() {
		defer close(done)

		switch len(channels) {
		case 0:
			return
		case 1:
			for val := range channels[0] {
				done <- val
			}
		default:
			switches := make([]chan interface{}, len(channels))
			for i := range switches {
				switches[i] = make(chan interface{})
				go func(ch <-chan interface{}, sw chan<- interface{}) {
					for val := range ch {
						sw <- val
					}
				}(channels[i], switches[i])
			}
			select {
			case <-switches[0]:
			case <-switches[1]:
			case <-switches[2]:
			case <-switches[3]:
			case <-switches[4]:
				// Обработка до 5 каналов. Можно расширить, если необходимо.
			}
		}
	}()

	return done
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}
