package arsenic

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Referer", func() {
		g.It("Should perform a basic query string request", func() {
			var opts requestOptions
			raw, _ := ioutil.ReadFile("./test_data/basic_request.json")
		    json.Unmarshal(raw, &opts)
			data, errs := DoRequest(opts)
			g.Assert(len(errs)).Equal(0)
			g.Assert(data).Equal("{\"monkey\":[\"flan\"]}\n")
		})
		g.It("Should perform a request with both a query string and query parameters", func() {
			var opts requestOptions
			raw, _ := ioutil.ReadFile("./test_data/qstring_qobj_request.json")
		    json.Unmarshal(raw, &opts)
			data, errs := DoRequest(opts)
			g.Assert(len(errs)).Equal(0)
			g.Assert(data).Equal("{\"monkey\":[\"flan\"]}\n")
		})
	})
}
