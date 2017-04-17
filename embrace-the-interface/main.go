// see https://www.youtube.com/watch?v=xyDkyFjzFVc
// see https://github.com/gophercon/2015-talks/blob/master/Tom%C3%A1s%20Senart%20-%20Embrace%20the%20Interface/ETI.pdf
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type Decorator func(Client) Client

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set(name, value)
			return c.Do(r)
		})
	}
}

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

func FaultTolerance(attempts int, backoff time.Duration) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			for i := 0; i < attempts; i++ {
				if res, err = c.Do(r); res.StatusCode == http.StatusOK && err == nil {
					break
				}
				time.Sleep(backoff * time.Duration(i))
			}
			return
		})
	}
}

func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func main() {
	c := Decorate(
		http.DefaultClient,
		Logging(log.New(os.Stdout, "", log.LstdFlags)),
		FaultTolerance(5, time.Second),
		Header("User-Agent", "Golang/1.1"),
	)
	r, _ := http.NewRequest("GET", "https://httpbin.org/status/200", nil)
	c.Do(r)
}
