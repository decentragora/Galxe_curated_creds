package pgn

import (
	"context"
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

// Check if address has sent 5 tx on public goods network
func CheckAddressHasSentFiveOrMoreTxOnPGN(clts *models.Clients) http.HandlerFunc {
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
		pgnClient := clts.PublicGoodsNetworkClient.Client

		ctx := context.Background()
		// PendingNonceAt returns the account nonce of the given account in the pending state.
		// This is the nonce that should be used for the next transaction.
		nonce, err := pgnClient.PendingNonceAt(ctx, commonAddress)
		if err != nil {
			log.Println("error getting nonce", err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		log.Println("nonce", nonce)

		if nonce >= 5 {
			resp := response{
				Result: true,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		} else {
			resp := response{
				Result: false,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		}
	}
}

func CheckAddressHasSentTenOrMoreTxOnPGN(clts *models.Clients) http.HandlerFunc {
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
		pgnClient := clts.PublicGoodsNetworkClient.Client

		ctx := context.Background()
		// PendingNonceAt returns the account nonce of the given account in the pending state.
		// This is the nonce that should be used for the next transaction.
		nonce, err := pgnClient.PendingNonceAt(ctx, commonAddress)
		if err != nil {
			log.Println("error getting nonce", err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		log.Println("nonce", nonce)

		if nonce >= 10 {
			resp := response{
				Result: true,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		} else {
			resp := response{
				Result: false,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		}
	}
}
