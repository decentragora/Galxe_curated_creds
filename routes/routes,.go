package routes

import (
	"dagora_galaxe_querys/models"
	_dagora_deployers "dagora_galaxe_querys/routes/dagoraDeployers"
	_dagora_membership "dagora_galaxe_querys/routes/dagoraMemberships"
	_dagora_dapp "dagora_galaxe_querys/routes/dagoradapp"
	_facet_check "dagora_galaxe_querys/routes/facet"
	_gitcoin_nfts "dagora_galaxe_querys/routes/gitcoinImpact"
	_impact_card "dagora_galaxe_querys/routes/impactCards"
	_pgn_check "dagora_galaxe_querys/routes/pgn"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//_logger.Info(`Incoming request to "` + req.RequestURI + `" from ` + req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}
func HandleRoutes(r *mux.Router, clts *models.Clients) {
	r.Use(loggingMiddleware)

	// Check if address has minted a dagora membership
	r.HandleFunc("/has_membership", _dagora_membership.CheckAddressHasMembership(clts))

	// Check if address has minted a dagora membership: ecclesia tier
	r.HandleFunc("/hold_ecclesia", _dagora_membership.CheckAddressHoldEcclesia(clts))

	// Check if address has minted a dagora membership: dagorian tier or higher
	r.HandleFunc("/hold_dagorian", _dagora_membership.CheckAddressHoldsDagorian(clts))

	// Check if address has minted a dagora membership: hoplite tier or higher
	r.HandleFunc("/hold_hoplite", _dagora_membership.CheckAddressHoldsHoplite(clts))

	// Check if address has minted a dagora membership: perclesian tier or higher
	r.HandleFunc("/hold_perclesia", _dagora_membership.CheckAddressHoldsPerclesia(clts))

	// Check if address has minted 1 of 3 gitcoin nfts
	r.HandleFunc("/has_atleast_one_gitcoin_nft", _gitcoin_nfts.CheckAddressHasAtleastOneGitcoinNFT(clts))

	// Check if address has minted at least 1 impact card nft
	r.HandleFunc("/has_minted_impact_cards", _impact_card.CheckAddressMintedImpactCards(clts))

	// Check if address has minted 2 or more impact card nfts
	r.HandleFunc("/has_two_or_more_impact_cards", _impact_card.CheckAddressForTwoOrMoreImpactCards(clts))

	// Check if address has minted a full season set of impact card nfts ie one of each id
	r.HandleFunc("/has_full_set_impact_cards", _impact_card.CheckAddressForFullSetOfImpactCards(clts))

	// Check if address has sent 5 or more tx on public goods network
	r.HandleFunc("/has_sent_five_or_more_tx_on_pgn", _pgn_check.CheckAddressHasSentFiveOrMoreTxOnPGN(clts))

	// Check if address has sent 10 or more tx on public goods network
	r.HandleFunc("/has_sent_ten_or_more_tx_on_pgn", _pgn_check.CheckAddressHasSentTenOrMoreTxOnPGN(clts))

	// Checks to see if a address has signed the manifesto
	r.HandleFunc("/has_signed_manifesto", _dagora_dapp.CheckAddressHasSignedManifesto(clts))

	// Check if address has donated to dagora facet
	r.HandleFunc("/has_donated_to_facet", _facet_check.CheckAddressDonatedToFacet(clts))

	// Check if address has deployed a contract using the dagora factory contracts on optimism
	r.HandleFunc("/has_deployed_factory_contract", _dagora_deployers.CheckAddressHasDeployedFactoryContract(clts))

}
