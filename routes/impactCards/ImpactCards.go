package impactcards

import (
	impactcardabi "dagora_galaxe_querys/contractabi/ImpactCardAbi"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressMintedImpactCards(clts *models.Clients) http.HandlerFunc {
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
		contractAddress := contracts.Contracts["ImpactCards"].Address
		optimsimClient := clts.OptimismClient.Client
		impactCardInstance, err := impactcardabi.NewImpactcardsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		var wg sync.WaitGroup
		totalBalance := big.NewInt(0)
		balanceChan := make(chan *big.Int, 15)

		for i := 1; i <= 15; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				balance, err := impactCardInstance.BalanceOf(&bind.CallOpts{}, commonAddress, big.NewInt(int64(i)))
				if err != nil {
					log.Println(err)
					helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
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

		// Balance is greater than 0 return true, else return false
		if totalBalance.Cmp(big.NewInt(0)) > 0 {
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

func CheckAddressForTwoOrMoreImpactCards(clts *models.Clients) http.HandlerFunc {
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
		contractAddress := contracts.Contracts["ImpactCards"].Address
		optimsimClient := clts.OptimismClient.Client
		impactCardInstance, err := impactcardabi.NewImpactcardsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		var wg sync.WaitGroup
		totalBalance := big.NewInt(0)
		balanceChan := make(chan *big.Int, 15)

		for i := 1; i <= 15; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				balance, err := impactCardInstance.BalanceOf(&bind.CallOpts{}, commonAddress, big.NewInt(int64(i)))
				log.Println("balance", balance)
				if err != nil {
					log.Println(err)
					helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
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

		// Balance is greater or equal to 2 return true, else return false
		if totalBalance.Cmp(big.NewInt(2)) >= 0 {
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

func CheckAddressForFullSetOfImpactCards(clts *models.Clients) http.HandlerFunc {
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
		contractAddress := contracts.Contracts["ImpactCards"].Address
		optimsimClient := clts.OptimismClient.Client
		impactCardInstance, err := impactcardabi.NewImpactcardsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		var wg sync.WaitGroup
		hasFullSet := true
		balanceChan := make(chan bool, 15)

		for i := 1; i <= 15; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				balance, err := impactCardInstance.BalanceOf(&bind.CallOpts{}, commonAddress, big.NewInt(int64(i)))
				if err != nil {
					log.Println(err)
					balanceChan <- false
					return
				}
				balanceChan <- balance.Cmp(big.NewInt(0)) > 0
			}(i)
		}

		go func() {
			wg.Wait()
			close(balanceChan)
		}()

		for hasBalance := range balanceChan {
			if !hasBalance {
				hasFullSet = false
				break
			}
		}

		resp := models.DefaultResponse{
			Result: hasFullSet,
		}
		helpers.RespondWithJSON(w, http.StatusOK, resp)

	}
}
