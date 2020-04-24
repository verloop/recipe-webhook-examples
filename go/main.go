package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/verloop/recipe-webhook-examples/go/common"
	"github.com/verloop/recipe-webhook-examples/go/webhookrequest"
	"github.com/verloop/recipe-webhook-examples/go/webhookresponse"
)

func main() {
	http.HandleFunc("/show_plans", common.Log(common.Auth(showplans)))
	http.HandleFunc("/get_plans", common.Log(common.Auth(getPlans)))
	http.HandleFunc("/get_payment_link", common.Log(common.Auth(getPaymentLink)))
	http.HandleFunc("/check_payment", common.Log(common.Auth(checkPayment)))
	http.ListenAndServe(":3000", nil)
}

func showplans(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var reqbody webhookrequest.WebhookRequest
	json.Unmarshal(body, &reqbody)
	json.NewEncoder(w).Encode(*webhookresponse.ConstructOperators())
}

func getPlans(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var reqbody webhookrequest.WebhookRequest
	json.Unmarshal(body, &reqbody)
	json.NewEncoder(w).Encode(*webhookresponse.ConstructPlansForOperator(&reqbody))
}

func getPaymentLink(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var reqbody webhookrequest.WebhookRequest
	json.Unmarshal(body, &reqbody)
	json.NewEncoder(w).Encode(*webhookresponse.ConstructPaymentLink(&reqbody))
}

func checkPayment(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var reqbody webhookrequest.WebhookRequest
	json.Unmarshal(body, &reqbody)
	json.NewEncoder(w).Encode(*webhookresponse.CheckPayment(&reqbody))
}
