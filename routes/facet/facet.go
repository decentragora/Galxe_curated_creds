package facet

import (
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	//"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressDonatedToFacet(clts *models.Clients) http.HandlerFunc {
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
		facetAddress := common.HexToAddress("0x9892ceD5B0a93673dDe5ee4cffD2F732fA1F406F")

		key := os.Getenv("OPSCAN_API_KEY")
		url := fmt.Sprintf("https://api-optimistic.etherscan.io/api?module=account&action=txlist&address=%s&apikey=%s", facetAddress, key)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		defer resp.Body.Close()

		var etherscanResponse models.EtherscanResponse
		err = json.NewDecoder(resp.Body).Decode(&etherscanResponse)
		if err != nil {
			log.Println(err)
			helpers.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		var wg sync.WaitGroup
		hasDonatedChan := make(chan bool)

		for _, tx := range etherscanResponse.Result {
			wg.Add(1)
			go func(tx models.Transaction) {
				defer wg.Done()
				if strings.ToLower(tx.From) == strings.ToLower(address) && strings.ToLower(tx.To) == strings.ToLower(facetAddress.String()) {
					hasDonatedChan <- true
				}
			}(tx)
		}

		// Close channel once all goroutines are done
		go func() {
			wg.Wait()
			close(hasDonatedChan)
		}()

		// Aggregate results from channel
		hasDonated := false
		for donation := range hasDonatedChan {
			if donation {
				hasDonated = true
				break
			}
		}

		// Respond with the result
		donatedResponse := models.DefaultResponse{Result: hasDonated}
		helpers.RespondWithJSON(w, http.StatusOK, donatedResponse)
	}
}
