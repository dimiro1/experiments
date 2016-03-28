package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "9090"), hystrixStreamHandler)

	for {
		select {
		case <-ticker.C:
			err := hystrix.Do("Localhost 8080", func() error {
				resp, err := http.Get("http://localhost:8080/")

				if resp != nil {
					resp.Body.Close()
				}

				return err
			}, nil)

			if err != nil {
				log.Printf("Error: %v", err)
			}
		}
	}
}
