package loadbalancing

import (
	"app-lb-round-robin/config"
)

type RoundRobin struct {
	current int
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{current: 0 }
}

func (r *RoundRobin) GetService() string{
	endpoint := config.ServicesList[r.current % len(config.ServicesList)]
	r.current++
	return endpoint
}




