package router

import (
	"fmt"
	"net/http"
)

func HandleUserRouter(res http.ResponseWriter, req *http.Request) (string, error) {
	method := req.Method
	url := req.URL
	path := url.Path

	// login
	if method == "POST" && path == "/api/user/login" {
		return `{"msg":"this is user login"}`, nil
	}

	return "", fmt.Errorf("blog router process failed")
}
