package arsenic

import (
	"errors"
	"encoding/json"
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

func toJson(obj map[string]string) (s string) {
	a, _ := json.Marshal(obj)
	s = string(a)
	return
}

func DoRequest(opts requestOptions) (data interface{}, errs []error) {
	r := request.CustomMethod(opts.Method, opts.Url).
		SetDebug(true).
		Set("User-Agent", "Super-spiffy arsenic golang useragent")
	if opts.QueryString != "" {
		r.Query(opts.QueryString)
	}
	if len(opts.QueryObj) != 0 {
		r.Query(toJson(opts.QueryObj))
	}
	resp, b, e := r.End()
	data = b
	if e != nil {
		errs = e
	}
	if resp.Status != "200 OK" {
		errs = append(errs, errors.New("non 200 status code: " + resp.Status))
	}
	return
}
