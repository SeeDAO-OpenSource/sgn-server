package server

import (
	"github.com/gin-gonic/gin"
	"github.com/waite-lee/sgn/pkg/app"
)

type ServerBuiler struct {
	AppBuilder *app.AppBuilder
	initors    []ServerConfigureFunc
	Options    *ServerOptions
}

func AddServer(builder *app.AppBuilder, options *ServerOptions) *ServerBuiler {
	builder.BindOptions("Server", options)
	return &ServerBuiler{
		AppBuilder: builder,
		initors:    make([]ServerConfigureFunc, 0),
		Options:    options,
	}
}

func (b *ServerBuiler) Configure(action ServerConfigureFunc) *ServerBuiler {
	b.initors = append(b.initors, action)
	return b
}

func (b *ServerBuiler) Build() (*Server, error) {
	g := gin.Default()
	server := NewServer(g, b.Options)
	for _, action := range b.initors {
		err := action(server)
		if err != nil {
			return nil, err
		}
	}
	return server, nil
}
