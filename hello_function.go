package helloworld

import (
	"encoding/json"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHttpFromGolang", HelloHTTP)
}

type HelloWorldResp struct {
	Result string
}


// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	res := HelloWorldResp{
		Result: "sukses dari function",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	raw, _ := json.Marshal(res)
	_, _ = w.Write(raw)
}