package main

import (
	proto "github.com/micro/examples/function/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// create a new function
	fnc := micro.NewFunction(
		micro.Name("go.micro.fnc.greeter"),
	)

	// init the command line
	fnc.Init()

	// register a handler
	fnc.Handle(new(Greeter))

	// run the function
	fnc.Run()
}
