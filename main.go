package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/parnurzeal/gorequest"
)

type requestOptions struct {
	Url string `json:url`
	Method string `json:method`
	Query string `json:query`
}

func RequestOptions(url, method, query string) (retval requestOptions) {
	retval = requestOptions{url, method, query}
	return
}

var request = gorequest.New()

func DoRequest(opts requestOptions) (data interface{}, err error) {
	r := request.CustomMethod(opts.Method, opts.Url).
		SetDebug(true).
		Set("User-Agent", "Super-spiffy golang useragent /u/cacahootie")
	if opts.Query != "" {
		r.Query(opts.Query)
	}
	resp, body, errs :=	r.End()
	data = body
	if errs != nil {
		fmt.Println(errs)
	} else if resp.Status != "200 OK" {
		fmt.Println(resp.Status)
		fmt.Println("non 200 status code")
	}
	return
}

func main() {
	var opts requestOptions
	raw, err := ioutil.ReadFile("./demo/config.json")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    json.Unmarshal(raw, &opts)
	//opts = RequestOptions("http://localhost:8080", "GET", "monkey=flan")
	DoRequest(opts)
}
