// Package micro is a runtime for the go.micro.run service
package micro

import (
	"errors"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-run"
	proto "github.com/micro/go-run/proto"

	"golang.org/x/net/context"
)

type microRuntime struct {
	Client proto.RuntimeClient
}

func (m *microRuntime) Fetch(url string, opts ...run.FetchOption) (*run.Source, error) {
	var options run.FetchOptions
	for _, o := range opts {
		o(&options)
	}

	rsp, err := m.Client.Fetch(context.Background(), &proto.FetchRequest{
		Url:    url,
		Update: options.Update,
	})
	if err != nil {
		return nil, err
	}

	return &run.Source{
		URL: rsp.Source.Url,
		Dir: rsp.Source.Dir,
	}, nil
}

func (m *microRuntime) Build(src *run.Source) (*run.Binary, error) {
	rsp, err := m.Client.Build(context.Background(), &proto.BuildRequest{
		Source: &proto.Source{
			Url: src.URL,
			Dir: src.Dir,
		},
	})
	if err != nil {
		return nil, err
	}

	return &run.Binary{
		Path: rsp.Binary.Path,
		Source: &run.Source{
			URL: rsp.Binary.Source.Url,
			Dir: rsp.Binary.Source.Dir,
		},
	}, nil
}

func (m *microRuntime) Exec(bin *run.Binary) (*run.Process, error) {
	rsp, err := m.Client.Exec(context.Background(), &proto.ExecRequest{
		Binary: &proto.Binary{
			Path: bin.Path,
			Source: &proto.Source{
				Url: bin.Source.URL,
				Dir: bin.Source.Dir,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &run.Process{
		ID: rsp.Process.Id,
		Binary: &run.Binary{
			Path: rsp.Process.Binary.Path,
			Source: &run.Source{
				URL: rsp.Process.Binary.Source.Url,
				Dir: rsp.Process.Binary.Source.Dir,
			},
		},
	}, nil
}

func (m *microRuntime) Kill(proc *run.Process) error {
	_, err := m.Client.Kill(context.Background(), &proto.KillRequest{
		Process: &proto.Process{
			Id: proc.ID,
			Binary: &proto.Binary{
				Path: proc.Binary.Path,
				Source: &proto.Source{
					Url: proc.Binary.Source.URL,
					Dir: proc.Binary.Source.Dir,
				},
			},
		},
	})
	return err
}

func (m *microRuntime) Wait(proc *run.Process) error {
	stream, err := m.Client.Wait(context.Background(), &proto.WaitRequest{
		Process: &proto.Process{
			Id: proc.ID,
			Binary: &proto.Binary{
				Path: proc.Binary.Path,
				Source: &proto.Source{
					Url: proc.Binary.Source.URL,
					Dir: proc.Binary.Source.Dir,
				},
			},
		},
	})
	if err != nil {
		return err
	}

	rsp, err := stream.Recv()
	if err != nil {
		return err
	}

	if len(rsp.Error) > 0 {
		return errors.New(rsp.Error)
	}

	return nil
}

func NewRuntime() run.Runtime {
	return &microRuntime{
		// TODO: make configurable
		Client: proto.NewRuntimeClient("go.micro.run", client.DefaultClient),
	}
}
