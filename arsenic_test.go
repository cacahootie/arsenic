package arsenic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	var opts requestOptions
	raw, err := ioutil.ReadFile("./demo/config.json")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    json.Unmarshal(raw, &opts)
	DoRequest(opts)
}
