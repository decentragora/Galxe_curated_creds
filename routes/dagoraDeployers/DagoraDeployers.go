package dagoradeployers

import (
	_dagora_erc20_factory "dagora_galaxe_querys/contractabi/dagoraFactorys/DagoraERC20Factory"

	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type body struct {
	Address string `json:"address"`
}

type response struct {
	Result bool `json:"result"`
}

// Check if address has Deployed one contract using the dagora factory contracts on optimism
// there are 6 factory contracts
func CheckAddressHasDeployedFactoryContract(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var b body
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

		hasDeployed := checkDagoraSimpleNFTFactory(commonAddress, clts)
		if err != nil {
			log.Println("error", err)
		}

		log.Println("has Deployed", hasDeployed)

	}
}

// function to check the dagora simple nft a deployer contrac tto see if a user has deployed a contract
func checkDagoraSimpleNFTFactory(address common.Address, clts *models.Clients) bool {
	optimismClient := clts.OptimismClient.Client
	contractAddress := contracts.Contracts["DagoraSimpleNFTFactory"].Address

	instance, err := _dagora_erc20_factory.NewDagoraerc20factory(contractAddress, optimismClient)
	if err != nil {
		return false
	}

	deployBal, err := instance.GetUserContracts(nil, address)
	if err != nil {
		return false
	}

	log.Println("deployBal", deployBal)
	return false
}

func checkDagoraERC20Factory(address common.Address) bool {
	// check if address has deployed a contract
	// get the total supply of the contract
	// get the balance of the address
	// if balance > 0 then return true
	// else return false
	return false
}

func checkDagoraPaymentSplitterFactory(address common.Address) bool {
	// check if address has deployed a contract
	// get the total supply of the contract
	// get the balance of the address
	// if balance > 0 then return true
	// else return false
	return false
}

func checkDagoraNFTAPlusFactory(address common.Address) bool {

	return false
}

func checkDagoraPowerNFTFactory(address common.Address) bool {

	return false
}

func checkDagoraPowerPlusNFTFactory(address common.Address) bool {

	return false
}
