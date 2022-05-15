package server

import (
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
	if s.Options.Port == 0 {
		return s.G.Run(":4100")
	}
	return s.G.Run(":" + strconv.Itoa(s.Options.Port))
}
