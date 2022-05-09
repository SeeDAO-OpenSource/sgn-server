package apiserver

import (
	"log"
	"strconv"

	"github.com/google/wire"
	blobv1 "github.com/waite-lee/sgn/internal/apiserver/blob/v1"
	nftv1 "github.com/waite-lee/sgn/internal/apiserver/nft/v1"
	"github.com/waite-lee/sgn/pkg/app"
	"github.com/waite-lee/sgn/pkg/server"
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
	//as.server.GEngine.Run(":5000")
	return as.server.GEngine.RunTLS(":"+strconv.Itoa(5000), "app/sgn_chain.crt", "app/sgn_key.key")
}

// func tlsHandler(port int) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		secureMiddleware := secure.New(secure.Options{
// 			SSLRedirect: true,
// 			SSLHost:     ":" + strconv.Itoa(port),
// 		})
// 		err := secureMiddleware.Process(c.Writer, c.Request)

// 		// If there was an error, do not continue.
// 		if err != nil {
// 			return
// 		}

// 		c.Next()
// 	}
// }
