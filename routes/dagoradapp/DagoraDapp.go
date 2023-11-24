package dagoradapp

import (
	_attest_station "dagora_galaxe_querys/contractabi/AttestStation"
	"dagora_galaxe_querys/contracts"
	"dagora_galaxe_querys/helpers"
	"dagora_galaxe_querys/models"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

func CheckAddressHasSignedManifesto(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var b models.DefaultBody
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		address := b.Address
		isValid := helpers.IsValidAddress(address)
		if isValid != true {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
			return
		}

		commonAddress := common.HexToAddress(address)
		atteststationAddress := contracts.Contracts["Attest Station"].Address
		attestSubjectAddress := common.HexToAddress("0x1De380594dE7ABA6442D879713c86Ba7395abE7B")
		attestKeyString := "6461676f72612e6d616e69666573746f2e616c69676e00000000000000000000"
		attestKey := common.HexToHash("0x" + attestKeyString)
		OptimismClient := clts.OptimismClient.Client

		attestationsInstance, err := _attest_station.NewAtteststation(atteststationAddress, OptimismClient)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Error creating contract instance")
			return
		}

		// params to pass to contract
		// address - users address
		// attestSubject - address of the subject of the attestation
		// attestKey - key of the attestation
		// check if address has signed manifesto
		response, err := attestationsInstance.Attestations(nil, commonAddress, attestSubjectAddress, attestKey)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Error creating contract instance")
			return
		}

		responseHex := hex.EncodeToString(response)
		// prepend with 0x
		responseHex = "0x" + responseHex

		if responseHex == "0x594145" {
			response := models.DefaultResponse{
				Result: true,
			}
			helpers.RespondWithJSON(w, http.StatusOK, response)
			return
		} else {
			response := models.DefaultResponse{
				Result: false,
			}
			helpers.RespondWithJSON(w, http.StatusOK, response)
			return
		}
	}
}

func CheckAddressHasMadeAttestation(clts *models.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			helpers.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var b models.DefaultBody
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		address := b.Address
		isValid := helpers.IsValidAddress(address)
		if isValid != true {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid address")
			return
		}

		// commonAddress := common.HexToAddress(address)
		// atteststationAddress := contracts.Contracts["Attest Station"].Address
		// OptimismClient := clts.OptimismClient.Client

		// check attestations contract event logs for address
	}
}
