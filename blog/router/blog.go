package router

import (
	"fmt"
	"learngo/blog/controller"
	"learngo/blog/model"
	"net/http"

	"github.com/gpmgo/gopm/modules/log"
)

func HandleBlogRouter(res http.ResponseWriter, req *http.Request) (string, error) {
	method := req.Method
	url := req.URL
	path := url.Path

	// 获取博客列表
	if method == "GET" && path == "/api/blog/list" {
		author := req.Form.Get("author")
		keyword := req.Form.Get("keyword")
		d, err := controller.GetList(author, keyword)
		if err != nil {
			log.Error("get list failed, err: %v", err)
			return model.ErrorModel("get blog list failed").String(), nil
		}
		return model.SuccessModel(d).String(), nil
	}

	if method == "GET" && path == "/api/blog/detail" {
		return `{"msg":"this is blog detail"}`, nil
	}

	if method == "POST" && path == "/api/blog/new" {
		return `{"msg":"this is blog new"}`, nil
	}

	if method == "POST" && path == "/api/blog/update" {
		return `{"msg":"this is blog update"}`, nil
	}

	if method == "POST" && path == "/api/blog/delete" {
		return `{"msg":"this is blog delete"}`, nil
	}

	return "", fmt.Errorf("blog router process failed")
}
