package main

import (
	"net/http"
	"time"

	webcache "github.com/iphuket/web-cache"
)

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.ListenAndServe(":5645", nil)
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
