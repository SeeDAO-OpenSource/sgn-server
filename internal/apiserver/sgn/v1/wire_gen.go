// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package sgnv1

import (
	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/internal/common"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob/file"
	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
	"github.com/SeeDAO-OpenSource/sgn/pkg/eth"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/SeeDAO-OpenSource/sgn/pkg/mvc"
)

// Injectors from wire_inject.go:

func BuildSgnServiceV1() (*SgnService, error) {
	etherScanOptions := _wireEtherScanOptionsValue
	httpClientOptions := _wireHttpClientOptionsValue
	client, err := erc721.GetClient(etherScanOptions, httpClientOptions)
	if err != nil {
		return nil, err
	}
	ethOptions := _wireEthOptionsValue
	ethclientClient, err := eth.GetClient(ethOptions)
	if err != nil {
		return nil, err
	}
	erc721Service := erc721.NewErc721Service(client, etherScanOptions, ethclientClient)
	requestClient := mvc.NewRequestClient(httpClientOptions)
	ipfsOptions := _wireIpfsOptionsValue
	ipfsClient, err := ipfs.GetClient(requestClient, ipfsOptions)
	if err != nil {
		return nil, err
	}
	fileBlobStoreOptions := _wireFileBlobStoreOptionsValue
	blobStore := file.NewFileBlobStore(fileBlobStoreOptions)
	mongoOptions := _wireMongoOptionsValue
	mongoClient, err := mongodb.GetClient(mongoOptions)
	if err != nil {
		return nil, err
	}
	erc721TransferLogRepo := NewMongoErc721TransferLogRepo(mongoClient)
	sgnTokenRepo := NewMongoDbSgnTokenRepo(mongoClient)
	sgnService := &SgnService{
		Erc:        erc721Service,
		Client:     ethclientClient,
		IpfsClient: ipfsClient,
		Blobstore:  blobStore,
		LogRepo:    erc721TransferLogRepo,
		TokenRepo:  sgnTokenRepo,
	}
	return sgnService, nil
}

var (
	_wireEtherScanOptionsValue     = EsOptions
	_wireHttpClientOptionsValue    = common.HttpOptions
	_wireEthOptionsValue           = common.EthOptions
	_wireIpfsOptionsValue          = common.IpfsOptions
	_wireFileBlobStoreOptionsValue = common.FileOptions
	_wireMongoOptionsValue         = common.MongoOptions
)
