package server

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServerOptions struct {
	Port      int
	ServerURL string
}

type Server struct {
	G       *gin.Engine
	Options *ServerOptions
}

func NewServer(g *gin.Engine, options *ServerOptions) *Server {
	return &Server{
		G:       g,
		Options: options,
	}
}

func (s Server) Run() error {
	addr := ":" + strconv.Itoa(s.Options.Port)
	if s.Options.Port == 0 {
		addr = ":4100"
	}
	log.Printf("**********GIN Listening: %s ***********\n", addr)
	return s.G.Run(addr)
}
