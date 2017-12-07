package main

import (
	"fmt"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-grpc"
	hello "github.com/micro/go-grpc/examples/greeter/server/proto/hello"
	"github.com/micro/go-micro/metadata"

	"golang.org/x/net/context"
)

var (
	serviceName string
)

func main() {
	service := grpc.NewService()
	service.Init(
		micro.Flags(cli.StringFlag{
			Name: "service_name",
			Value: "go.micro.srv.greeter",
			Destination: &serviceName,
		}),
	)

	// use the generated client stub
	cl := hello.NewSayClient(serviceName, service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
