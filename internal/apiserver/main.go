package apiserver

import (
	"github.com/google/wire"
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
	return ApiServer{
		server:     server,
		appContext: context,
	}
}

func (as *ApiServer) Run() error {
	nftv1.InstallNftV1(as.appContext, as.server)
	as.server.Init()
	return as.server.GEngine.Run(":5000")
}
