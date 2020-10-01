package main

import (
	"app-lb-round-robin/gateway"
	"app-lb-round-robin/loadbalancing"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	TEXT = "text"
)

func main() {
	//Initialize the dependencies
	var lb = loadbalancing.NewRoundRobin()
	var proxy = gateway.NewProxy()
	var gw = gateway.NewGateway(lb, proxy)


	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/*proxyPath", gw.ReverseProxy())
	r.POST("/*proxyPath", gw.ReverseProxy())

	r.Run("0.0.0.0:5000")
}
