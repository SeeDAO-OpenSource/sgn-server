package server

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/gin-gonic/gin"
)

type ServerBuiler struct {
	App     *app.AppBuilder
	initors []ServerConfigureFunc
	Options *ServerOptions
}

func NewServerBuilder(builder *app.AppBuilder) *ServerBuiler {
	return &ServerBuiler{
		App:     builder,
		initors: make([]ServerConfigureFunc, 0),
		Options: &ServerOptions{},
	}
}

func (b *ServerBuiler) Configure(action ServerConfigureFunc) *ServerBuiler {
	b.initors = append(b.initors, action)
	return b
}

func (b *ServerBuiler) Add(action func(*ServerBuiler) error) *ServerBuiler {
	action(b)
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
