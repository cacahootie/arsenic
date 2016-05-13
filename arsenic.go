package arsenic

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type requestOptions struct {
	Url string `json:url`
	Method string `json:method`
	QueryString string `json:queryString`
	QueryObj map[string]string `json:queryObj`
}

func RequestOptions(url, method, queryString string) (retval requestOptions) {
	retval.Url = url
	retval.Method = method
	retval.QueryString = queryString
	return
}

var request = gorequest.New()

func DoRequest(opts requestOptions) (data interface{}, err error) {
	r := request.CustomMethod(opts.Method, opts.Url).
		SetDebug(true).
		Set("User-Agent", "Super-spiffy arsenic golang useragent")
	if opts.QueryString != "" {
		r.Query(opts.QueryString)
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
