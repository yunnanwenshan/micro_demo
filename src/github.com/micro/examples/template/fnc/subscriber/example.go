package subscriber

import (
	"github.com/micro/go-log"

	example "github.com/micro/examples/template/fnc/proto/example"
	"golang.org/x/net/context"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
