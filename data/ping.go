package data

import "github.com/google/uuid"

type ping struct {
	Id     string `json:"id"`
	Method string `json:"method"`
}

type pingResponse struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Result struct {
	} `json:"result"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		Count         int    `json:"count"`
	} `json:"rateLimits"`
}

// Create default `ping` message to be sent to websocket
func messagePing() ping {
	id := uuid.New()

	return ping{
		Id:     id.String(),
		Method: "ping",
	}
}
