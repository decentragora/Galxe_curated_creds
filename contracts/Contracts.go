package contracts

import (
	//"dagora_galaxe_querys/clients"
	"dagora_galaxe_querys/models"

	"github.com/ethereum/go-ethereum/common"
)

var NFTContracts = map[string]models.NFTContract{

	"Gitcoin Web3 Grants Report 01: PDF Onchain": {
		Address: common.HexToAddress("0x9c0b5f09042b4bc9c5CE9eb75150c701976D7de8"),
		Network: "pgn",
	},
	"Impact on Zora Network": {
		Address: common.HexToAddress("0x81D226fb36cA785583E79E84312335d0e166D59B"),
		Network: "Zora",
	},
	"Gitcoin Impact Report": {
		Address: common.HexToAddress("0x50296fA66EA1486e3255F2F1f4a815A58F2A5fB2"),
		Network: "mainnet",
	},
	"ImpactCards": {
		Address: common.HexToAddress("0x0D3061bfF9be308E38da4699d44c6e9892E57E60"),
		Network: "Optimism",
	},
	"Dagora Membership": {
		Address: common.HexToAddress("0x9839C0f5715BDA074af1C6802cb9D157169c18D5"),
		Network: "Optimism",
	},
}

var CampaginTwoContracts = map[string]models.NFTContract{
	"Gitcoin Web3 Grants Report 01: PDF Onchain": {
		Address: common.HexToAddress("0x9c0b5f09042b4bc9c5CE9eb75150c701976D7de8"),
		Network: "pgn",
	},
	"ImpactOnZoraNetwork": {
		Address: common.HexToAddress("0x81D226fb36cA785583E79E84312335d0e166D59B"),
		Network: "Zora",
	},
	"Gitcoin Impact Report": {
		Address: common.HexToAddress("0x50296fA66EA1486e3255F2F1f4a815A58F2A5fB2"),
		Network: "mainnet",
	},
	"ImpactCards": {
		Address: common.HexToAddress("0x0D3061bfF9be308E38da4699d44c6e9892E57E60"),
		Network: "Optimism",
	},
}
