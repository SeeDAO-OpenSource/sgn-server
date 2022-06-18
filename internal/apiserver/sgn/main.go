package sgn

import (
	"errors"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddSgn(builder *server.ServerBuiler) error {
	erc721.AddErc721Services(builder.App)
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	builder.App.ConfigureServices(func() error {
		services.AddTransient(createSgnService)
		services.AddTransient(func(c *services.Container) *Erc721TransferLogRepo {
			client := services.Get[mongo.Client]()
			if client == nil {
				return nil
			}
			repo := NewMongoErc721TransferLogRepo(client)
			return &repo
		})
		services.AddTransient(func(c *services.Container) *SgnTokenRepo {
			client := services.Get[mongo.Client]()
			if client == nil {
				return nil
			}
			repo := NewMongoDbSgnTokenRepo(client)
			return &repo
		})
		return nil
	})
	return nil
}

func initRoute(g *gin.Engine) error {
	sgnCtl := newSgnController()
	route(&sgnCtl, g)
	SubscribeTransferLogs()
	return nil
}

func SubscribeTransferLogs() error {
	sgnService := services.Get[SgnService]()
	if sgnService == nil {
		return errors.New("sgn service is nil")
	}
	addresses, err := sgnService.GetExistsAddresses()
	if err != nil {
		return err
	}
	if len(addresses) > 0 {
		err = sgnService.SubscribeTransferLogs(addresses)
		if err != nil {
			return err
		}
	}
	return nil
}

func createSgnService(c *services.Container) *SgnService {
	srv := &SgnService{}
	srv.Client = services.Get[ethclient.Client]()
	srv.Erc = services.Get[erc721.Erc721Service]()
	srv.IpfsClient = services.Get[ipfs.IpfsClient]()
	srv.LogRepo = *services.Get[Erc721TransferLogRepo]()
	srv.TokenRepo = *services.Get[SgnTokenRepo]()
	srv.Blobstore = *services.Get[blob.BlobStore]()
	return srv
}
