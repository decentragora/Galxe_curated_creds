package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MainnetClient struct {
	Client *ethclient.Client
}

type OptimismClient struct {
	Client *ethclient.Client
}

type ZoraClient struct {
	Client *ethclient.Client
}

type PublicGoodsNetworkClient struct {
	Client *ethclient.Client
}

type Clients struct {
	MainnetClient            *MainnetClient
	OptimismClient           *OptimismClient
	ZoraClient               *ZoraClient
	PublicGoodsNetworkClient *PublicGoodsNetworkClient
}

type NFTContract struct {
	Address common.Address
	Network string
}
