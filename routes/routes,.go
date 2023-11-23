package routes

import (
	"dagora_galaxe_querys/models"
	_campaign_two "dagora_galaxe_querys/routes/campaign_two"
	_dagora_membership "dagora_galaxe_querys/routes/dagoraMemberships"
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

	r.HandleFunc("/funding_public_goods", _campaign_two.CheckAddressForFundingPublicGoods(clts))

	r.HandleFunc("/bonus_campaing_altra_card_holder", _campaign_two.CheckAddressForBonusCampaignAltraCardHolder(clts)).Methods("GET")

	r.HandleFunc("/has_membership", _dagora_membership.CheckAddressHasMembership(clts))

	r.HandleFunc("/hold_ecclesia", _dagora_membership.CheckAddressHoldEcclesia(clts)).Methods("GET")

	r.HandleFunc("/hold_hoplite", _dagora_membership.CheckAddressHoldsHoplite(clts)).Methods("GET")

	r.HandleFunc("/hold_perclesia", _dagora_membership.CheckAddressHoldsPerclesia(clts)).Methods("GET")
}
