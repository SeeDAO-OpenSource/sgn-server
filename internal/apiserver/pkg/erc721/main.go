package erc721

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc721"
	"github.com/nanmu42/etherscan-api"
)

type Erc721Service struct {
	options   *EtherScanOptions
	client    *etherscan.Client
	ethClient *ethclient.Client
}

func NewErc721Service(client *etherscan.Client, options *EtherScanOptions, ethClient *ethclient.Client) *Erc721Service {
	service := &Erc721Service{
		options:   options,
		client:    client,
		ethClient: ethClient,
	}
	return service
}

func (srv *Erc721Service) GetTransferLogs(address *string, page int, pageSize int) ([]etherscan.ERC721Transfer, error) {
	return srv.client.ERC721Transfers(address, nil, nil, nil, page, pageSize, true)
}

func (srv *Erc721Service) GetToken(ethClient *ethclient.Client, address *string, tokenId *big.Int) (*TokenInfo, error) {
	contrctAddr := common.HexToAddress(*address)
	token, err := erc721.NewErc721(contrctAddr, ethClient)
	if err != nil {
		return nil, err
	}
	info := TokenInfo{
		TokenId: tokenId.Int64(),
	}
	name, err := token.Name(nil)
	info.Name = name
	tokenURI, err := token.TokenURI(nil, tokenId)
	info.TokenURI = tokenURI
	info.Contract = *address
	tokenAddress, err := token.OwnerOf(nil, tokenId)
	info.Owner = tokenAddress.Hex()
	info.ID = info.Contract + strconv.FormatInt(info.TokenId, 10)
	return &info, nil
}

func (srv *Erc721Service) SubscribeFilterLogs(addresses []string, callback func(log *types.Log)) error {
	ethAddresses := make([]common.Address, len(addresses))
	for i, address := range addresses {
		ethAddresses[i] = common.HexToAddress(address)
	}
	tranferEventSignature := []byte("Transfer(address,address,uint256)")
	hash := crypto.Keccak256Hash(tranferEventSignature)
	query := ethereum.FilterQuery{
		Addresses: ethAddresses,
		Topics:    [][]common.Hash{{hash}},
	}
	logs := make(chan types.Log)
	sub, err := srv.ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err == nil {
		//defer sub.Unsubscribe()
		go func() {
			for {
				select {
				case err = <-sub.Err():
					log.Fatal(err)
				case vLog := <-logs:
					fmt.Println(vLog)
					callback(&vLog)
				}
			}
		}()
	}
	return err
}

func (srv *Erc721Service) GetTransferLog(log *types.Log) (*etherscan.ERC721Transfer, error) {
	tranferLog := etherscan.ERC721Transfer{
		Hash:             log.TxHash.Hex(),
		From:             log.Topics[1].Hex(),
		To:               log.Topics[2].Hex(),
		BlockNumber:      int(log.BlockNumber),
		BlockHash:        log.BlockHash.Hex(),
		ContractAddress:  log.Address.Hex(),
		TokenID:          (*etherscan.BigInt)(log.Topics[3].Big()),
		TransactionIndex: int(log.TxIndex),
	}
	block, err := srv.ethClient.BlockByHash(context.Background(), log.BlockHash)
	if err != nil {
		return nil, err
	}
	tx, isppendding, err := srv.ethClient.TransactionByHash(context.Background(), log.TxHash)
	if err != nil {
		return nil, err
	} else if isppendding {
		err = fmt.Errorf("tx is pending")
	}
	tranferLog.Nonce = int(tx.Nonce())
	tranferLog.GasPrice = (*etherscan.BigInt)(tx.GasPrice())
	tranferLog.GasUsed = int(tx.Gas())
	tranferLog.CumulativeGasUsed = int(tx.GasFeeCap().Int64())
	tranferLog.Input = string(tx.Data())
	tranferLog.TimeStamp = etherscan.Time(time.Unix(int64(block.Time()), 0))
	return &tranferLog, err
}
