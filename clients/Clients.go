package clients

import (
	"dagora_galaxe_querys/models"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rpc"
)

/*
Clients Needed
	-Mainnet
	-Optimism
	-Public Goods Network


pgn rpc: https://rpc.publicgoods.network

*/

type ClientWrapper struct {
	*models.Clients
}

func InitClients() (*models.Clients, error) {
	mainnetUrl := fmt.Sprintf("https://mainnet.infura.io/v3/%s", os.Getenv("INFURAID"))
	optimsimUrl := fmt.Sprintf("https://optimism-mainnet.infura.io/v3/%s", os.Getenv("INFURAID"))
	zoraUrl := "https://rpc.zora.energy"
	pgnUrl := "https://rpc.publicgoods.network"

	mainnetClient, err := ethclient.Dial(mainnetUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	optimismClient, err := ethclient.Dial(optimsimUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	zoraClient, err := ethclient.Dial(zoraUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pgnClient, err := ethclient.Dial(pgnUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &models.Clients{
		MainnetClient: &models.MainnetClient{
			Client: mainnetClient,
		},
		OptimismClient: &models.OptimismClient{
			Client: optimismClient,
		},
		ZoraClient: &models.ZoraClient{
			Client: zoraClient,
		},
		PublicGoodsNetworkClient: &models.PublicGoodsNetworkClient{
			Client: pgnClient,
		},
	}, nil
}

func (cw *ClientWrapper) GetClient(network []string) map[string]*ethclient.Client {
	clientMap := make(map[string]*ethclient.Client)

	for _, net := range network {
		switch net {
		case "mainnet":
			clientMap["mainnet"] = cw.MainnetClient.Client
		case "optimism":
			clientMap["optimism"] = cw.OptimismClient.Client
		case "zora":
			clientMap["zora"] = cw.ZoraClient.Client
		case "pgn":
			clientMap["pgn"] = cw.PublicGoodsNetworkClient.Client
		}
	}

	return clientMap
}
