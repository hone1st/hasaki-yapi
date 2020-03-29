package util

import (
	"errors"
	"github.com/nahid/gohttp"
)

type Request struct {
	Url   string                 `json:"url"`
	Query map[string]interface{} `json:"query"`
}

// post请求
func (r *Request) Post() (*gohttp.Response, error) {
	req := gohttp.NewRequest()
	if r.Query != nil {
		req = req.JSON(r.Query)
	}
	resp, err := req.Post(r.Url)
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode() != 200 {
		return nil, errors.New("请求无法响应！")
	}
	return resp, err
}

// get请求
func (r *Request) Get() (*gohttp.Response, error) {
	req := gohttp.NewRequest()
	if r.Query != nil {
		req = req.JSON(r.Query)
	}
	resp, err := req.Get(r.Url)
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode() != 200 {
		return nil, errors.New("请求无法响应！")
	}
	return resp, err
}
