# SeeDao 官网v2 NFT 后端服务

[![Build](https://github.com/SeeDAO-OpenSource/nft-server/actions/workflows/build.yml/badge.svg?branch=main&event=status)](https://github.com/SeeDAO-OpenSource/nft-server/actions/workflows/build.yml) [![Release](https://github.com/SeeDAO-OpenSource/nft-server/actions/workflows/release.yml/badge.svg?branch=main&event=release)](https://github.com/SeeDAO-OpenSource/nft-server/actions/workflows/release.yml)


提供官网中以太坊相关的服务接口

- 鲜花榜
- ...


## 鲜花榜

记录领取过 SeeDao NFT 账户的信息，提供展示最新的拥有 SeeDAO V2 NFT 的账户以及metadata的数据访问接口.

提供的接口如下：

- [x] 获取拥有 SeeDAO V2 NFT 的发放信息， 按获取时间倒序排序，返回最新的部分数据
- [ ] ...

[查看Demo页面](http://124.221.160.98:5000/app/demo)

## 引用的开源库

- [gin](https://github.com/gin-gonic/gin)
- [wire](https://github.com/google/wire)
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [mongo-driver](https://github.com/mongodb/mongo-go-driver)
- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [etherscan-api](https://github.com/nanmu42/etherscan-api)
- [resty](https://github.com/go-resty/resty)