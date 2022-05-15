package blobv1

import (
	"github.com/waite-lee/sgn/pkg/server"
)

func AddBlobV1(buider *server.ServerBuiler) error {
	buider.Configure(initRoute)
	return nil
}

func initRoute(s *server.Server) error {
	nftCtl := newBlobController()
	route(&nftCtl, s.G)
	return nil
}
