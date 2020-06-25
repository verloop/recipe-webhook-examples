package server

import (
	"fmt"
	"net/http"
)

func (api *RechargeAPI) showplans(w http.ResponseWriter, r *http.Request) {
	var reqbody WebhookRequest
	if err := api.decode(w, r, &reqbody); err != nil {
		api.respond(w, r, "", http.StatusBadRequest)
		return
	}
	api.respond(w, r, api.ConstructOperators(), http.StatusOK)
}

func (api *RechargeAPI) getPlans(w http.ResponseWriter, r *http.Request) {
	var reqbody WebhookRequest
	if err := api.decode(w, r, &reqbody); err != nil {
		api.respond(w, r, "", http.StatusBadRequest)
		return
	}
	api.respond(w, r, api.ConstructPlansForOperator(&reqbody), http.StatusOK)
}

func (api *RechargeAPI) getPaymentLink(w http.ResponseWriter, r *http.Request) {
	var reqbody WebhookRequest
	if err := api.decode(w, r, &reqbody); err != nil {
		fmt.Println(err)
		api.respond(w, r, "", http.StatusBadRequest)
		return
	}
	api.respond(w, r, api.ConstructPaymentLink(&reqbody), http.StatusOK)
}

func (api *RechargeAPI) checkPayment(w http.ResponseWriter, r *http.Request) {
	var reqbody WebhookRequest
	err := api.decode(w, r, &reqbody)
	if err != nil {
		api.respond(w, r, "", http.StatusBadRequest)
		return
	}
	api.respond(w, r, api.CheckPayment(&reqbody), http.StatusOK)
}
