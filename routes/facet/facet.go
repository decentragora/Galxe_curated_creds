// package facet

// import (
// 	"context"
// 	"dagora_galaxe_querys/helpers"
// 	"dagora_galaxe_querys/models"
// 	"math/big"
// 	"net/http"

// 	"github.com/ethereum/go-ethereum/common"
// )

// func CheckAddressDonatedToFacet(clts *models.Clients) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet && r.Method != http.MethodPost {
// 			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
// 			return
// 		}

// 		address := r.URL.Query().Get("address")
// 		isValid := helpers.IsValidAddress(address)
// 		if isValid != true {
// 			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
// 			return
// 		}
// 		commonAddress := common.HexToAddress(address)
// 		optimsimClient := clts.OptimismClient.Client
// 		facetAddress := common.HexToAddress("0x9892ceD5B0a93673dDe5ee4cffD2F732fA1F406F")

// 		// use etherscan api to check transactions and check from and to addresses

// 		// No donation transaction found
// 		helpers.RespondWithJSON(w, http.StatusOK, false)

// 	}
// }
