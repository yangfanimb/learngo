package main

import (
	"learngo/blog/router"
	"net/http"
)

const errMsg = `404 Not Found`

func serverHandle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "Application/json")

	// parse query and post data
	req.ParseForm()

	var err error
	var msg string
	msg, err = router.HandleBlogRouter(res, req)
	if err == nil {
		res.Write([]byte(msg))
		return
	}

	msg, err = router.HandleUserRouter(res, req)
	if err == nil {
		res.Write([]byte(msg))
		return
	}

	res.WriteHeader(404)
	res.Header().Set("Content-type", "text/")
	res.Write([]byte(errMsg))
	return
}

func main() {
	http.HandleFunc("/", serverHandle)
	panic(http.ListenAndServe(":8881", nil))
}
