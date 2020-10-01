package gateway

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {

}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (proxy Proxy) ReverseProxy(c *gin.Context, target string) {
	director := func(req *http.Request) {
		r := c.Request
		req = r
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.Header["my-header"] = []string{r.Header.Get("my-header")}
		// Golang camelcases headers
		delete(req.Header, "My-Header")
	}
	reverseProxy := &httputil.ReverseProxy{Director: director}
	reverseProxy.ServeHTTP(c.Writer, c.Request)
}