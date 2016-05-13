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
			raw, _ := ioutil.ReadFile("./demo/basic_request.json")	    
		    json.Unmarshal(raw, &opts)
			DoRequest(opts)
		})
	})
}
