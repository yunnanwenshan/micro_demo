package main

import (
	"fmt"

	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"

	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService()
	service.Init()

	cl := hello.NewSayClient("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"User": "john",
		"ID":   "1",
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
