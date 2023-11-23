package dagoramemberships

import (
	dagoramemberships "dagora_galaxe_querys/contractabi/dagoraMembershipAbi"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"math/big"

	// "math/big"
	// "sync"
	"log"
	"net/http"

	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressHasMembership(clts *models.Clients) http.HandlerFunc {
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
		contractAddress := contracts.NFTContracts["Dagora Membership"].Address
		optimsimClient := clts.OptimismClient.Client

		membershipInstance, err := dagoramemberships.NewDagoramembershipsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		balance, err := membershipInstance.BalanceOf(nil, commonAddress)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if balance.Cmp(big.NewInt(0)) == 0 {
			helpers.RespondWithJSON(w, http.StatusOK, false)
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, true)
		return
	}
}
