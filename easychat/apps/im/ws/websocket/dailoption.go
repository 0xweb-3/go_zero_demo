package websocket

import (
	"net/http"
)

type DailOptions func(option *dailOption)

type dailOption struct {
	patten string
	header http.Header
}

func newDailOptions(opts ...DailOptions) dailOption {
	o := dailOption{
		patten: "/ws",
		header: nil,
	}

	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithClientPatten(patten string) DailOptions {
	return func(opt *dailOption) {
		opt.patten = patten
	}
}

func WithClientHeader(header http.Header) DailOptions {
	return func(opt *dailOption) {
		opt.header = header
	}
}
