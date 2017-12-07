package main

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	hello "github.com/micro/examples/greeter/srv/proto/hello"

	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	// Use the generated client stub
	cl := hello.NewSayClient("go.micro.srv.greeter", client.DefaultClient)

	// Make request
	rsp, err := cl.Hello(context.Background(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
