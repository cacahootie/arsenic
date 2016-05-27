# Arsenic

Arsenic is an http request client for go which focuses on enabling declarative queries, i.e. requests which are defined by a template JSON and combined with a context JSON, and which returns the result and errors from the query.

# Example

The request json:
```javascript
{
	"url":"http://localhost:8080",
	"method":"GET",
	"queryString":"monkey=flan"
}
```

The invoking code:
```go
var opts requestOptions
raw, _ := ioutil.ReadFile("./test_data/qstring_qobj_post.json")
json.Unmarshal(raw, &opts)
data, errs := DoRequest(opts)
```
