package blobv1

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
)

func AddBlob(buider *server.ServerBuiler) error {
	buider.Configure(initRoute)
	return nil
}

func initRoute(s *server.Server) error {
	sgnCtl := newBlobController()
	route(&sgnCtl, s.G)
	return nil
}
