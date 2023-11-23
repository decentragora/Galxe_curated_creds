package campaign_two

import (
	"context"
	impactcardabi "dagora_galaxe_querys/contractabi/ImpactCardAbi"
	zora1155impl "dagora_galaxe_querys/contractabi/Zora1155Impl"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"log"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressForFundingPublicGoods(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		address := r.URL.Query().Get("address")
		isValid := helpers.IsValidAddress(address)
		if isValid != true {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
			return
		}
		commonAddress := common.HexToAddress(address)

		impactOnZora := make(chan bool)
		gitcoinImpactReport := make(chan bool)
		gitcoinWeb3GrantsReport := make(chan bool)
		moreThan5TxOnPublicGoodsNetwork := make(chan bool)
		moreThan5TxOnOptimism := make(chan bool)

		go func() {
			impactOnZora <- checkAddressForImpactOnZoraNetwork(commonAddress, clts)
		}()

		go func() {
			gitcoinImpactReport <- checkAddressForGitcoinImpactReport(commonAddress, clts)
		}()

		go func() {
			gitcoinWeb3GrantsReport <- checkAddressForGitcoinWeb3GrantsReport(commonAddress, clts)
		}()

		go func() {
			moreThan5TxOnPublicGoodsNetwork <- checkAddressSent5TxOnPublicGoodsNetwork(commonAddress, clts)
		}()

		go func() {
			moreThan5TxOnOptimism <- checkAddressSent5TxOnOptimism(commonAddress, clts)
		}()

		impactOnZoraResult := <-impactOnZora
		gitcoinImpactReportResult := <-gitcoinImpactReport
		gitcoinWeb3GrantsReportResult := <-gitcoinWeb3GrantsReport
		moreThan5TxOnPublicGoodsNetworkResult := <-moreThan5TxOnPublicGoodsNetwork
		moreThan5TxOnOptimismResult := <-moreThan5TxOnOptimism

		// in order to qualify for the campaign, the address must have:
		// atleast 1 of the 3 nfts in their wallet
		// AND
		// have sent 5 tx on public goods network
		// OR
		// have sent 5 tx on optimism
		if (impactOnZoraResult || gitcoinImpactReportResult || gitcoinWeb3GrantsReportResult) && (moreThan5TxOnPublicGoodsNetworkResult || moreThan5TxOnOptimismResult) {
			helpers.RespondWithJSON(w, http.StatusOK, true)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, false)
		return
	}
}

func CheckAddressForBonusCampaignAltraCardHolder(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		address := r.URL.Query().Get("address")
		isValid := helpers.IsValidAddress(address)
		if isValid != true {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
			return
		}
		commonAddress := common.HexToAddress(address)

		hasTwoOrMoreImpactCards := checkAddressForTwoOrMoreImpactCards(commonAddress, clts)
		if hasTwoOrMoreImpactCards {
			helpers.RespondWithJSON(w, http.StatusOK, true)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, false)
		return
	}
}

func checkAddressForTwoOrMoreImpactCards(address common.Address, clts *models.Clients) bool {
	optimsimClient := clts.OptimismClient.Client
	contractAddress := contracts.CampaginTwoContracts["ImpactCards"].Address

	contractInstance, err := impactcardabi.NewImpactcardsabi(contractAddress, optimsimClient)
	if err != nil {
		log.Println("error creating contract instance", err)
		return false
	}

	var wg sync.WaitGroup
	totalBalance := big.NewInt(0)
	balanceChan := make(chan *big.Int, 15)

	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			balance, err := contractInstance.BalanceOf(&bind.CallOpts{}, address, big.NewInt(int64(i)))
			if err != nil {
				log.Println("error getting balance", err)
				return
			}
			balanceChan <- balance
		}(i)
	}

	go func() {
		wg.Wait()
		close(balanceChan)
	}()

	for balance := range balanceChan {
		totalBalance.Add(totalBalance, balance)
	}

	if totalBalance.Cmp(big.NewInt(2)) >= 0 {
		return true
	}
	return false
}

// Contract is on Zora network
// use zora1155impl to get the contract, balanceOf address, and tokenId 1
func checkAddressForImpactOnZoraNetwork(address common.Address, clts *models.Clients) bool {
	zoraClient := clts.ZoraClient.Client
	contractAddress := contracts.CampaginTwoContracts["ImpactOnZoraNetwork"].Address

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
	contractAddress := contracts.CampaginTwoContracts["Gitcoin Impact Report"].Address

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
	contractAddress := contracts.CampaginTwoContracts["Gitcoin Web3 Grants Report 01: PDF Onchain"].Address

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

// Check if address has sent 5 tx on public goods network
func checkAddressSent5TxOnPublicGoodsNetwork(address common.Address, clts *models.Clients) bool {
	pgnClient := clts.PublicGoodsNetworkClient.Client

	ctx := context.Background()
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	nonce, err := pgnClient.PendingNonceAt(ctx, address)
	if err != nil {
		log.Println("error getting nonce", err)
		return false
	}

	if nonce-1 >= 5 {
		return true
	}
	return false
}

func checkAddressSent5TxOnOptimism(address common.Address, clts *models.Clients) bool {
	optimismClient := clts.OptimismClient.Client

	ctx := context.Background()

	nonce, err := optimismClient.PendingNonceAt(ctx, address)
	if err != nil {
		log.Println("error getting nonce", err)
		return false
	}

	if nonce-1 >= 5 {
		return true
	}
	return false
}
