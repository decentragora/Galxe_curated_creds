package gitcoinimpact

import (
	zora1155impl "dagora_galaxe_querys/contractabi/Zora1155Impl"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressHasAtleastOneGitcoinNFT(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var b models.DefaultBody
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid body")
			return
		}

		address := b.Address
		isValid := helpers.IsValidAddress(address)
		if isValid != true {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
			return
		}

		commonAddress := common.HexToAddress(address)
		impactOnZora := make(chan bool)
		gitcoinImpactReport := make(chan bool)
		gitcoinWeb3GrantsReport := make(chan bool)

		go func() {
			impactOnZora <- checkAddressForImpactOnZoraNetwork(commonAddress, clts)
		}()

		go func() {
			gitcoinImpactReport <- checkAddressForGitcoinImpactReport(commonAddress, clts)
		}()

		go func() {
			gitcoinWeb3GrantsReport <- checkAddressForGitcoinWeb3GrantsReport(commonAddress, clts)
		}()

		impactOnZoraResult := <-impactOnZora
		gitcoinImpactReportResult := <-gitcoinImpactReport
		gitcoinWeb3GrantsReportResult := <-gitcoinWeb3GrantsReport

		// if any of the results are true, return true
		if impactOnZoraResult || gitcoinImpactReportResult || gitcoinWeb3GrantsReportResult {
			resp := models.DefaultResponse{
				Result: true,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		}

		resp := models.DefaultResponse{
			Result: false,
		}
		helpers.RespondWithJSON(w, http.StatusOK, resp)
		return
	}
}

// Contract is on Zora network
// use zora1155impl to get the contract, balanceOf address, and tokenId 1
func checkAddressForImpactOnZoraNetwork(address common.Address, clts *models.Clients) bool {
	zoraClient := clts.ZoraClient.Client
	contractAddress := contracts.Contracts["ImpactOnZoraNetwork"].Address

	contractInstance, err := zora1155impl.NewZora1155impl(contractAddress, zoraClient)
	if err != nil {
		log.Println("error creating contract instance", err)
		return false
	}

	balance, err := contractInstance.BalanceOf(&bind.CallOpts{}, address, big.NewInt(1))
	if err != nil {
		log.Println("error getting balance", err)
		return false
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		return false
	}
	return true
}

// Contract is on Mainnet
// use zora1155impl to get the contract, balanceOf address, and tokenId 1
func checkAddressForGitcoinImpactReport(address common.Address, clts *models.Clients) bool {
	mainnetClient := clts.MainnetClient.Client
	contractAddress := contracts.Contracts["Gitcoin Impact Report"].Address

	contractInstance, err := zora1155impl.NewZora1155impl(contractAddress, mainnetClient)
	if err != nil {
		log.Println("error creating contract instance", err)
		return false
	}

	balance, err := contractInstance.BalanceOf(&bind.CallOpts{}, address, big.NewInt(1))
	if err != nil {
		log.Println("error getting balance", err)
		return false
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		return false
	}
	return true
}

// Contract is on Public Goods Network
// use zora1155impl to get the contract, balanceOf address, and tokenId 1
func checkAddressForGitcoinWeb3GrantsReport(address common.Address, clts *models.Clients) bool {
	pgnClient := clts.PublicGoodsNetworkClient.Client
	contractAddress := contracts.Contracts["Gitcoin Web3 Grants Report 01: PDF Onchain"].Address

	contractInstance, err := zora1155impl.NewZora1155impl(contractAddress, pgnClient)
	if err != nil {
		log.Println("error creating contract instance", err)
		return false
	}

	balance, err := contractInstance.BalanceOf(&bind.CallOpts{}, address, big.NewInt(1))
	if err != nil {
		log.Println("error getting balance", err)
		return false
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		return false
	}
	return true
}
