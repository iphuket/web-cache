package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	webcache "github.com/iphuket/web-cache"
)

func main() {
	app := gin.New()
	app.Any("/get", getGin)
	app.Any("/set", setGin)
	app.Run(":5645")
}
func getGin(c *gin.Context) {
	get(c.Writer, c.Request)
}

func setGin(c *gin.Context) {
	set(c.Writer, c.Request)
}
func get(rw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	if len(key) < 1 {
		rw.Write([]byte("error"))
		return
	}
	str, bool := webcache.Get(key)
	if bool {
		rw.Write([]byte(str))
		return
	}
	rw.Write([]byte("error"))
}
func set(rw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	if len(key) < 1 {
		rw.Write([]byte("error"))
		return
	}
	value := req.FormValue("value")
	if len(value) < 1 {
		rw.Write([]byte("error"))
		return
	}
	webcache.Set(key, value, 110*time.Minute)
	rw.Write([]byte("success"))
}
