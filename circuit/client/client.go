package main

import (
	"log"
	"time"

	circuit "github.com/rubyist/circuitbreaker"
)

func main() {
	client := circuit.NewHTTPClient(time.Second*5, 1, nil)
	ticker := time.NewTicker(1 * time.Second)
	state := client.Panel.Subscribe()

	go func() {
		for {
			select {
			case ev := <-state:
				var name string

				switch ev.Event {
				case circuit.BreakerTripped:
					name = "BreakerTripped"
				case circuit.BreakerReset:
					name = "BreakerReset"
				case circuit.BreakerFail:
					name = "BreakerFail"
				case circuit.BreakerReady:
					name = "BreakerReady"
				}

				log.Printf("Event: %s", name)
			}
		}
	}()

	for {
		select {
		case <-ticker.C:
			_, err := client.Get("http://localhost:8080")

			log.Printf("Error: %v", err)
		}
	}
}
