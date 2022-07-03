package sgn

import (
	"errors"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/di"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/server"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SgnApi(builder *server.ServerBuiler) error {
	erc721.AddErc721Services(builder.App)
	builder.Configure(func(s *server.Server) error {
		return initRoute(s.G)
	})
	builder.App.ConfigureServices(func() error {
		di.AddTransient(createSgnService)
		di.AddTransient(func(c *di.Container) *Erc721TransferLogRepo {
			client := di.Get[mongo.Client]()
			if client == nil {
				return nil
			}
			repo := NewMongoErc721TransferLogRepo(client)
			return &repo
		})
		di.AddTransient(func(c *di.Container) *SgnTokenRepo {
			client := di.Get[mongo.Client]()
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
	sgnService := di.Get[SgnService]()
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

func createSgnService(c *di.Container) *SgnService {
	srv := &SgnService{}
	srv.Client = di.Get[ethclient.Client]()
	srv.Erc = di.Get[erc721.Erc721Service]()
	srv.IpfsClient = di.Get[ipfs.IpfsClient]()
	srv.LogRepo = *di.Get[Erc721TransferLogRepo]()
	srv.TokenRepo = *di.Get[SgnTokenRepo]()
	srv.Blobstore = *di.Get[blob.BlobStore]()
	return srv
}
