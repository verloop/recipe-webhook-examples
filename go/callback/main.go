package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/verloop/recipe-webhook-examples/go/server"
)

func main() {
	url := "https://dinesh.stage.verloop.io/webhooks/v1/1592491518.hkd06zh3vzdx"
	res := server.ConstructOperatorsForTrigger()
	// var req webhookrequest.WebhookRequest
	// req.State = make(map[string]string)
	// req.State["order_id"] = "hsjhfhsdfjks"
	// success := webhookresponse.CheckPayment(&req)
	jres, _ := json.Marshal(res)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jres))
	if err != nil {
		fmt.Errorf(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
