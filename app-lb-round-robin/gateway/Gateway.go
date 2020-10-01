package gateway

import (
	"app-lb-round-robin/loadbalancing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Gateway struct {
	servicePicker loadbalancing.IServicePicker
	proxy *Proxy
}

func NewGateway(servicePicker loadbalancing.IServicePicker, proxy *Proxy) *Gateway {
	return &Gateway{servicePicker: servicePicker, proxy: proxy}
}

func (gateway Gateway) ForwardRequest(c *gin.Context, path string)  {
	endpoint := gateway.servicePicker.GetService()
	gateway.proxy.ReverseProxy(c, endpoint+ path)
}

func (gateway Gateway) ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := gateway.servicePicker.GetService()
		remote, _ := url.Parse(target)

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}

	/*target := "localhost:4000"

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			r := c.Request
			req = r
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Header["my-header"] = []string{r.Header.Get("my-header")}
			// Golang camelcases headers
			delete(req.Header, "My-Header")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}*/
}
