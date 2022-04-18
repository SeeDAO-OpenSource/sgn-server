package apiserver

import (
	"log"

	"github.com/google/wire"
	blobv1 "github.com/waite-lee/nftserver/internal/apiserver/blob/v1"
	nftv1 "github.com/waite-lee/nftserver/internal/apiserver/nft/v1"
	"github.com/waite-lee/nftserver/pkg/app"
	"github.com/waite-lee/nftserver/pkg/server"
)

var ApiServerSet = wire.NewSet(
	server.NewServerContext,
	NewApiServer,
	wire.Value(AsOptions),
)

type ApiServer struct {
	server     *server.ServerContext
	appContext *app.AppContext
}

func NewApiServer(server *server.ServerContext, context *app.AppContext) ApiServer {
	var as = ApiServer{
		server:     server,
		appContext: context,
	}
	err := nftv1.InstallNftV1(as.appContext, as.server)
	if err == nil {
		err = blobv1.InstallBlobV1(as.appContext, as.server)
	}
	if err != nil {
		log.Fatal(err)
	}
	return as
}

func (as *ApiServer) Run() error {
	as.server.Init()
	as.server.GEngine.Run(":5000")
	return nil
}
