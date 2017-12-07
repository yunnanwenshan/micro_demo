package handler

import (
	example "github.com/micro/examples/template/fnc/proto/example"
	"golang.org/x/net/context"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
