// The MIT License (MIT)

// Copyright (c) 2016 Claudemiro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Inspired by the springboot actuator health check

package main

import (
	"encoding/json"
	"net/http"
)

const (
	// Up Convenient constant value representing up state.
	Up = "up"

	// Down Convenient constant value representing down state.
	Down = "down"

	// OutOfService Convenient constant value representing out-of-service state.
	OutOfService = "outOfService"

	// Unknown Convenient constant value representing unknown state.
	Unknown = "unknown"
)

// Health is a health status struct
type Health struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info,omitempty"`
}

// NewHealth return a new Health with status Down
func NewHealth() Health {
	h := Health{}
	h.Down()

	return h
}

// IsUp returns true if Status is Up
func (h Health) IsUp() bool {
	return h.Status == Up
}

// IsDown returns true if Status is Down
func (h Health) IsDown() bool {
	return h.Status == Down
}

// IsOutOfService returns true if Status is IsOutOfService
func (h Health) IsOutOfService() bool {
	return h.Status == OutOfService
}

// Down set the status to Down
func (h *Health) Down() {
	h.Status = Down
}

// OutOfService set the status to OutOfService
func (h *Health) OutOfService() {
	h.Status = OutOfService
}

// Unknown set the status to Unknown
func (h *Health) Unknown() {
	h.Status = Unknown
}

// Up set the status to Up
func (h *Health) Up() {
	h.Status = Up
}

// HealthAggregator is a struct used to aggregate Health instances into a final one.
type HealthAggregator struct {
	Indicators map[string]HealthIndicator
}

// NewHealthAggregator used to aggregate Health instances into a final one.
func NewHealthAggregator() HealthAggregator {
	return HealthAggregator{
		Indicators: make(map[string]HealthIndicator),
	}
}

// AddHealthIndicator add a HealthIndicator to the aggregator
func (ha *HealthAggregator) AddHealthIndicator(name string, health HealthIndicator) {
	ha.Indicators[name] = health
}

// Health returns a composite Health
func (ha HealthAggregator) Health() Health {
	health := NewHealth()
	health.Up()

	healths := map[string]Health{}

	for name, indicator := range ha.Indicators {
		h, err := indicator.Health()

		if err != nil {
			health.Down()
		}

		if !health.IsDown() && h.IsDown() {
			health.Down()
		}

		healths[name] = h
	}

	health.Info = healths
	return health
}

func (ha HealthAggregator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := ha.Health()

	if health.IsDown() {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// CompositeHealthIndicator aggregate a list of HealthIndicator
type CompositeHealthIndicator struct {
	Indicators map[string]HealthIndicator
}

// NewCompositeHealthIndicator creates a new CompositeHealthIndicator
func NewCompositeHealthIndicator() CompositeHealthIndicator {
	return CompositeHealthIndicator{
		Indicators: make(map[string]HealthIndicator),
	}
}

// AddHealthIndicator add a HealthIndicator to the aggregator
func (i *CompositeHealthIndicator) AddHealthIndicator(name string, health HealthIndicator) {
	i.Indicators[name] = health
}

// Health returns the health check of CompositeHealthIndicator
func (i CompositeHealthIndicator) Health() (Health, error) {
	var err error

	health := NewHealth()
	health.Up()

	healths := map[string]Health{}

	for name, indicator := range i.Indicators {
		h, err := indicator.Health()

		if err != nil {
			health.Down()
		}

		if !health.IsDown() && h.IsDown() {
			health.Down()
		}

		healths[name] = h
	}

	health.Info = healths

	return health, err
}

// HealthIndicator is a interface used to provide an indication of application health.
type HealthIndicator interface {
	// Health Returns a Health object
	Health() (Health, error)
}

// URLHealthIndicatorResponse is a struct used to return the status of the retrieved url
type URLHealthIndicatorResponse struct {
	Status int `json:"status"`
}

// URLHealthIndicator holds the URL to check
type URLHealthIndicator struct {
	URL string
}

// Health returns the health check of URLHealthIndicator
func (g URLHealthIndicator) Health() (Health, error) {
	req, err := http.NewRequest("GET", g.URL, nil)

	health := NewHealth()
	health.Up()

	if err != nil {
		health.Down()
		health.Info = URLHealthIndicatorResponse{Status: http.StatusBadRequest}

		return health, err
	}

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		health.Down()
		health.Info = URLHealthIndicatorResponse{Status: http.StatusBadRequest}

		return health, err
	}

	if resp.StatusCode != http.StatusOK {
		health.Down()
	}

	health.Info = URLHealthIndicatorResponse{
		Status: resp.StatusCode,
	}

	return health, nil
}

func main() {
	indicators := NewCompositeHealthIndicator()
	indicators.AddHealthIndicator("GuiaBolso", URLHealthIndicator{URL: "http://www.guiabolso.com.br"})
	indicators.AddHealthIndicator("Google", URLHealthIndicator{URL: "http://www.google.com"})

	aggregator := NewHealthAggregator()
	aggregator.AddHealthIndicator("Websites", indicators)
	aggregator.AddHealthIndicator("G1", URLHealthIndicator{URL: "http://g1.globo.com/index.html"})

	http.Handle("/health/", aggregator)
	http.ListenAndServe(":8080", nil)
}
