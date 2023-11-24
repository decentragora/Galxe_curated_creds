package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DefaultBody struct {
	Address string `json:"address"`
}

type DefaultResponse struct {
	Result bool `json:"result"`
}

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

type EtherscanResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}
