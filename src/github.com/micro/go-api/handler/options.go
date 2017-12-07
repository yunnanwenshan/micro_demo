package handler

import (
	"github.com/micro/go-micro"
)

type Options struct {
	Namespace string
	Service   micro.Service
}

type Option func(o *Options)

// NewOptions fills in the blanks
func NewOptions(opts ...Option) Options {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	// create service if its blank
	if options.Service == nil {
		WithService(micro.NewService())(&options)
	}

	// set namespace if blank
	if len(options.Namespace) == 0 {
		WithNamespace("go.micro.api")(&options)
	}

	return options
}

func WithNamespace(s string) Option {
	return func(o *Options) {
		o.Namespace = s
	}
}

func WithService(s micro.Service) Option {
	return func(o *Options) {
		o.Service = s
	}
}
