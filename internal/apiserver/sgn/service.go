package sgn

import (
	"log"
	"math/big"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/pkg/blob"
	"github.com/SeeDAO-OpenSource/sgn/pkg/ipfs"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
)

type SgnService struct {
	Erc        *erc721.Erc721Service
	Client     *ethclient.Client
	IpfsClient *ipfs.IpfsClient
	Blobstore  blob.BlobStore
	LogRepo    Erc721TransferLogRepo
	TokenRepo  SgnTokenRepo
}

func (srv *SgnService) GetOwners(address string, skip int64, limit int64) ([]erc721.TokenInfo, error) {
	return srv.TokenRepo.GetList(address, skip, limit)
}

func (srv *SgnService) GetTransferLogs(contract string) ([]ERC721Transfer, error) {
	data, err := srv.Erc.GetTransferLogs(&contract, 0, 5)
	if err != nil {
		return nil, err
	}
	return convertTransfer(data), nil

}

func (srv *SgnService) PullData(contract *string, skip int, tokens []string, logging bool) error {
	data, err := srv.Erc.GetTransferLogs(contract, 0, 5)
	if err != nil {
		return err
	}
	logData := []etherscan.ERC721Transfer{}
	for i, v := range data {
		if i < skip {
			continue
		}

		if len(tokens) > 0 {
			exists := false
			for _, t := range tokens {
				exists = (t == v.TokenID.Int().String())
			}
			if !exists {
				continue
			}
		}
		logData = append(logData, v)
	}
	logPrintf(logging, "共查询到 %v 条事件日志\n", len(logData))
	return srv.pullTansferLogs(contract, logData, logging)
}

func (srv *SgnService) GetExistsAddresses() ([]string, error) {
	return srv.LogRepo.GetAddresses()
}

func (srv *SgnService) SubscribeTransferLogs(addresses []string) error {
	err1 := srv.Erc.SubscribeFilterLogs(addresses, func(eventlog *types.Log) {
		address := eventlog.Address.String()
		tranfer, err := srv.Erc.GetTransferLog(eventlog)
		if err == nil {
			srv.pullTansferLogs(&address, []etherscan.ERC721Transfer{*tranfer}, true)
		} else {
			log.Fatalln(err)
		}
	})
	return err1
}

func (srv *SgnService) GetTokenImage(token int64, address string, process *blob.Process) (*blob.BlobReader, error) {
	tokeInfo, err := srv.TokenRepo.Get(token, address)
	if err != nil {
		return nil, err
	}
	uri := &tokeInfo.Metadata.Image
	if !srv.Blobstore.Exists(uri) {
		if err := srv.SaveImage(uri, true); err != nil {
			return nil, err
		}
	}
	return srv.Blobstore.Read(uri, process)
}

// func (srv *SgnService) isSeedaoSgn(transfer *etherscan.ERC721Transfer) (bool, error) {
// 	tranx, _, err := srv.Client.TransactionByHash(context.Background(), common.HexToHash(transfer.Hash))
// 	if err != nil {
// 		return false, err
// 	}
// 	abi, err := abi.JSON(strings.NewReader(seedaoSgnAbi))
// 	if err != nil {
// 		return false, err
// 	}
// 	data := tranx.Data()
// 	method, err := abi.MethodById(data[:4])
// 	if err != nil {
// 		return false, err
// 	}
// 	return method.Name == "mintWhiteList", nil
// }

func (srv *SgnService) pullTansferLogs(contract *string, data []etherscan.ERC721Transfer, logging bool) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	// transfer日志写入数据库
	if err := srv.LogRepo.InsertMany(convertTransfer(data)); err != nil {
		return err
	}
	// 写入token信息到数据库
	return srv.pullTokens(data, logging, contract)
}

func (srv *SgnService) pullTokens(data []etherscan.ERC721Transfer, logging bool, contract *string) error {
	result := []*erc721.TokenInfo{}
	for _, v := range data {
		logPrintf(logging, "正在解析 TokenId: %v \n", v.TokenID.Int())
		tokenId := v.TokenID
		tokenInfo, err := srv.parseToken(contract, tokenId.Int(), logging)
		tokenInfo.TimeStamp = v.TimeStamp.Time().UnixMilli()
		tokenInfo.ID = v.Hash
		if err != nil {
			logPrintf(logging, "解析Token信息出错: %v\n", err.Error())
			return err
		}
		result = append(result, tokenInfo)
	}
	if err := srv.TokenRepo.InsertMany(result); err != nil {
		return err
	}
	logPrintf(logging, "插入Token信息成功\n")
	logPrintf(logging, "开始缓存图片...\n")
	errResult := []int64{}
	for _, v := range result {
		if v.Metadata.Image != "" {
			err := srv.SaveImage(&v.Metadata.Image, logging)
			if err != nil {
				logPrintf(logging, "保存图片: %s 出错, TokenId: %v\n", v.Metadata.Image, v.TokenId)
				errResult = append(errResult, v.TokenId)
			}
		}
	}

	logPrintf(logging, "缓存图片完成, 出错数量: %v\n", len(errResult))
	if len(errResult) > 0 {
		logPrintf(logging, "出错Tokens: %v", errResult)
	}
	return nil
}

func (srv *SgnService) parseToken(contract *string, tokenID *big.Int, logging bool) (*erc721.TokenInfo, error) {
	tokenInfo, err := srv.Erc.GetToken(srv.Client, contract, tokenID)
	if err != nil {
		return nil, err
	}
	uri := tokenInfo.TokenURI
	metadata, err := srv.parseMetadata(uri, logging)
	if err != nil {
		return nil, err
	}
	tokenInfo.Metadata = *metadata
	return tokenInfo, nil
}

func (srv *SgnService) parseMetadata(uri string, logging bool) (*erc721.TokenMetadata, error) {
	content, err := srv.IpfsClient.GetString(uri)
	if err != nil {
		logPrintf(logging, "获取Metadata内容出错: %v\n", err.Error())
		return nil, err
	}
	metadata := erc721.ParseMetadata(content)
	return &metadata, nil
}

func (srv *SgnService) SaveImage(imageUri *string, logging bool) error {
	if !srv.Blobstore.Exists(imageUri) {
		logPrintf(logging, "保存Image: %v\n", *imageUri)
		image, err := srv.IpfsClient.GetContent(*imageUri)
		if err != nil {
			logPrintf(logging, "获取Image内容出错: %v\n", err.Error())
			return err
		}
		err = srv.Blobstore.Save(imageUri, &image, false)
		if err != nil {
			logPrintf(logging, "保存Iamge出错: %v\n", err.Error())
			return err
		}
	} else {
		logPrintf(logging, "Image已存在: %v\n", *imageUri)
	}
	return nil
}

func logPrintf(islog bool, format string, v ...interface{}) {
	if islog {
		log.Printf(format, v...)
	}
}

func convertTransfer(data []etherscan.ERC721Transfer) []ERC721Transfer {
	result := []ERC721Transfer{}
	for _, v := range data {
		result = append(result, mapTransfer(&v))
	}
	return result
}

func mapTransfer(transfer *etherscan.ERC721Transfer) ERC721Transfer {
	result := ERC721Transfer{
		BlockNumber:       transfer.BlockNumber,
		TimeStamp:         transfer.TimeStamp.Time(),
		Hash:              transfer.Hash,
		Nonce:             transfer.Nonce,
		BlockHash:         transfer.BlockHash,
		From:              transfer.From,
		ContractAddress:   transfer.ContractAddress,
		To:                transfer.To,
		TokenID:           transfer.TokenID.Int().Int64(),
		TokenName:         transfer.TokenName,
		TokenSymbol:       transfer.TokenSymbol,
		TokenDecimal:      transfer.TokenDecimal,
		TransactionIndex:  transfer.TransactionIndex,
		Gas:               transfer.Gas,
		GasPrice:          transfer.GasPrice.Int().Int64(),
		GasUsed:           transfer.GasUsed,
		CumulativeGasUsed: transfer.CumulativeGasUsed,
		Input:             transfer.Input,
		Confirmations:     transfer.Confirmations,
	}
	result.ID = result.Hash
	return result
}
