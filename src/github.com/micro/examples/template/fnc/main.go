package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/examples/template/fnc/handler"
	"github.com/micro/examples/template/fnc/subscriber"
)

func main() {
	// New Service
	function := micro.NewFunction(
		micro.Name("go.micro.fnc.template"),
		micro.Version("latest"),
	)

	// Register Handler
	function.Handle(new(handler.Example))

	// Register Struct as Subscriber
	function.Subscribe("topic.go.micro.fnc.template", new(subscriber.Example))

	// Initialise function
	function.Init()

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
