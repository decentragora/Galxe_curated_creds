package dagoramemberships

import (
	dagoramemberships "dagora_galaxe_querys/contractabi/dagoraMembershipAbi"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type body struct {
	Address string `json:"address"`
}

type response struct {
	Result bool `json:"result"`
}

func CheckAddressHasMembership(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		//unmarshal body
		var b body
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid body")
			return
		}

		address := b.Address
		// GET ADDRESS from body

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
			resp := response{
				Result: false,
			}
			helpers.RespondWithJSON(w, http.StatusOK, resp)
			return
		}

		resp := response{
			Result: true,
		}
		helpers.RespondWithJSON(w, http.StatusOK, resp)
		return
	}
}

func CheckAddressHoldEcclesia(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
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
		contractAddress := contracts.NFTContracts["Dagora Membership"].Address
		optimsimClient := clts.OptimismClient.Client

		membershipInstance, err := dagoramemberships.NewDagoramembershipsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenId, err := membershipInstance.AddressTokenIds(nil, commonAddress)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if tokenId == big.NewInt(0) {
			helpers.RespondWithJSON(w, http.StatusOK, false)
			return
		}

		// check that tokenId is ecclesia tier
		tier, err := membershipInstance.GetMembershipTier(nil, tokenId)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// tiers are
		// 0 - ecclesia
		// 1 - dagora
		// 2 - hoplite
		// 3 - perclesian
		// if tier is 0, then return true
		if tier == 0 {
			helpers.RespondWithJSON(w, http.StatusOK, true)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, false)
		return
	}
}

func CheckAddressHoldsHoplite(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
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
		contractAddress := contracts.NFTContracts["Dagora Membership"].Address
		optimsimClient := clts.OptimismClient.Client

		membershipInstance, err := dagoramemberships.NewDagoramembershipsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenId, err := membershipInstance.AddressTokenIds(nil, commonAddress)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if tokenId.Cmp(big.NewInt(0)) == 0 {
			helpers.RespondWithJSON(w, http.StatusOK, false)
			return
		}

		// check that tokenId is hoplite tier
		tier, err := membershipInstance.GetMembershipTier(nil, tokenId)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// tiers are
		// 0 - ecclesia
		// 1 - dagora
		// 2 - hoplite
		// 3 - perclesian
		// if tier is 2, then return true
		if tier == 2 {
			helpers.RespondWithJSON(w, http.StatusOK, true)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, false)
		return
	}
}

func CheckAddressHoldsPerclesia(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
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
		contractAddress := contracts.NFTContracts["Dagora Membership"].Address
		optimsimClient := clts.OptimismClient.Client

		membershipInstance, err := dagoramemberships.NewDagoramembershipsabi(contractAddress, optimsimClient)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		tokenId, err := membershipInstance.AddressTokenIds(nil, commonAddress)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		if tokenId.Cmp(big.NewInt(0)) == 0 {
			helpers.RespondWithJSON(w, http.StatusOK, false)
			return
		}

		// check that tokenId is hoplite tier
		tier, err := membershipInstance.GetMembershipTier(nil, tokenId)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		// tiers are
		// 0 - ecclesia
		// 1 - dagora
		// 2 - hoplite
		// 3 - perclesian
		// if tier is 3, then return true
		if tier == 3 {
			helpers.RespondWithJSON(w, http.StatusOK, true)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, false)
		return
	}
}
