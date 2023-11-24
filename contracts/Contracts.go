package contracts

import (
	//"dagora_galaxe_querys/clients"
	"dagora_galaxe_querys/models"

	"github.com/ethereum/go-ethereum/common"
)

/*

DagoraERC20Factory  0x50B47F663a0645bb1feB6fF381579C6B74Ed6531
https://optimistic.etherscan.io/address/0x50b47f663a0645bb1feb6ff381579c6b74ed6531

DagoraSimpleNFTFactory:  0x9307807F3b6901bBDb458550158177f3B5Be16Da
https://optimistic.etherscan.io/address/0x9307807f3b6901bbdb458550158177f3b5be16da

DagoraPaymentSplitterFactory Factory: 0x0cBCb8120e648ba9F80Ac314c5d9fDe1cd2cFf76
https://optimistic.etherscan.io/address/0x0cbcb8120e648ba9f80ac314c5d9fde1cd2cff76

DagoraNFTAPlusFactory: 0xdc8E5beA9064e715f146120C7dd767af7f92CE90
https://optimistic.etherscan.io/address/0xdc8e5bea9064e715f146120c7dd767af7f92ce90

DagoraPowerNFTFactory  0x0214bCF488d29c867E5dC7c5BD19D934500527FA
https://optimistic.etherscan.io/address/0x0214bcf488d29c867e5dc7c5bd19d934500527fa

DagoraPowerPlusNFTFactory: 0xb3D0862F6030B654189372c0d099F89d115301e6
https://optimistic.etherscan.io/address/0xb3d0862f6030b654189372c0d099f89d115301e6

*/

var Contracts = map[string]models.NFTContract{

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
	"DagoraERC20Factory": {
		Address: common.HexToAddress("0x50B47F663a0645bb1feB6fF381579C6B74Ed6531"),
		Network: "Optimism",
	},
	"DagoraSimpleNFTFactory": {
		Address: common.HexToAddress("0x9307807F3b6901bBDb458550158177f3B5Be16Da"),
		Network: "Optimism",
	},
	"DagoraPaymentSplitterFactory": {
		Address: common.HexToAddress("0x0cBCb8120e648ba9F80Ac314c5d9fDe1cd2cFf76"),
		Network: "Optimism",
	},
	"DagoraNFTAPlusFactory": {
		Address: common.HexToAddress("0xdc8E5beA9064e715f146120C7dd767af7f92CE90"),
		Network: "Optimism",
	},
	"DagoraPowerNFTFactory": {
		Address: common.HexToAddress("0x0214bCF488d29c867E5dC7c5BD19D934500527FA"),
		Network: "Optimism",
	},
	"DagoraPowerPlusNFTFactory": {
		Address: common.HexToAddress("0xb3D0862F6030B654189372c0d099F89d115301e6"),
		Network: "Optimism",
	},
	"Attest Station": {
		Address: common.HexToAddress("0xEE36eaaD94d1Cc1d0eccaDb55C38bFfB6Be06C77"),
		Network: "Optimism",
	},
}
