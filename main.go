package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type requestOptions struct {
	url, method string
}

func RequestOptions (destUrl string) (retval requestOptions) {
	retval = requestOptions{"http://www.reddit.com/top.json", "GET"}
	return
}

var request = gorequest.New()

func DoRequest (opts requestOptions) {
	var (
		resp gorequest.Response
		body string
		errs []error
	)
	resp, body, errs = request.CustomMethod(opts.method, opts.url).
		SetDebug(true).
		Set("User-Agent","Super-spiffy golang useragent /u/cacahootie").
	End()
	_ = body
	if errs != nil {
		fmt.Println(errs)
		return
	} else if resp.Status != "200 OK" {
		fmt.Println(resp.Status)
		fmt.Println("non 200 status code")
		return
	} else {
		fmt.Print(body)
	}
	_ = resp
	_ = body
	_ = errs
}

func main() {
	request_opts := RequestOptions("http://www.reddit.com/top.json")
	DoRequest(request_opts)
}
