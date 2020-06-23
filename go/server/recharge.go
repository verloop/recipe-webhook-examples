package server

import (
	"fmt"
	"log"
)

//WebhookResponse has info about next block and quick replies
//ConstructOperators returns list of all operators
func (r *RechargeAPI) ConstructOperators() *WebhookResponse {
	res := NewWebhookResponse("Show_Operators")
	a1 := NewActions("Get_Plans_Wait")
	a1.AddVariable("operator", "Vodafone")

	a2 := NewActions("Get_Plans_Wait")
	a2.AddVariable("operator", "Airtel")

	o1 := NewOperators("Vodafone", a1)
	o1.AddTriggers("myvphone")
	o2 := NewOperators("Airtel", a2)
	o2.AddTriggers("myaphone")
	res.AddExportList("OperatorList", o1)
	res.AddExportList("OperatorList", o2)
	return &res
}

func ConstructOperatorsForTrigger() *WebhookResponse {
	res := NewWebhookResponse("Show_Operators")
	a1 := NewActions("Get_Plans_Wait")
	a1.AddVariable("operator", "Vodafone")

	a2 := NewActions("Get_Plans_Wait")
	a2.AddVariable("operator", "Airtel")
	o1 := NewOperators("Vodafone", a1)
	o1.AddTriggers("myvphone")
	o2 := NewOperators("Airtel", a2)
	o2.AddTriggers("myaphone")
	res.AddExportList("OperatorList", o1)
	res.AddExportList("OperatorList", o2)
	return &res
}

//ConstructPlansForOperator constructs a recharge plan
func (r *RechargeAPI) ConstructPlansForOperator(req *WebhookRequest) *WebhookResponse {

	operatorKey, ok := req.Variables["operator"]
	if !ok {
		log.Println("operator key not present")
		req := NewWebhookResponse("Invalid_Operator")
		return &req
	}

	operator, ok := operatorKey["parsed_value"]

	if !ok || !checkValidOperator(operator.(string)) {
		log.Println("invalid operator")
		req := NewWebhookResponse("Invalid_Operator")
		return &req
	}

	res := NewWebhookResponse("Show_Plans_Text")

	a := NewActions("Do_Recharge")
	a.AddVariable("amount", "100")
	button := NewButton("Select", "postback", a)
	url := NewURL("Know More", "web_url", "https://verloop.io")

	plan1 := NewPlan("Data 1 GB (28 days)", "Rs. 100", nil)
	plan1.AddButton(button, url)

	a = NewActions("Do_Recharge")
	a.AddVariable("amount", "150")
	button = NewButton("Select", "postback", a)

	plan2 := NewPlan("Full Talk time (84 days)", "Rs. 150", nil)
	plan2.AddButton(button, url)
	res.AddExportList("PlanList", plan1, plan2)
	return &res
}

//ConstructPaymentLink returns a dummy payment link
func (r *RechargeAPI) ConstructPaymentLink(req *WebhookRequest) *WebhookResponse {
	amountkey, ok := req.Variables["amount"]
	if !ok {
		res := NewWebhookResponse("Invalid_Options")
		return &res
	}

	amount, ok := amountkey["value"]
	if !ok {
		res := NewWebhookResponse("Invalid_Options")
		return &res
	}

	res := NewWebhookResponse("")
	res.AddState("order_id", "NXPAOMDAJDAY")

	button := NewButton("Payment Done", "postback", NewActions("Verify_Payment"))
	url := NewURL(fmt.Sprintf("Pay Rs : %s INR", amount.(string)), "web_url", "https://verloop.io/pricing.html")
	res.AddExportList("PaymentOptions", button, url)
	return &res
}

//CheckPayment checks for a order_id in the request state
//If order_id is present it returns Order_Success to be executed next else returns Order_Failure block to be executed next
func (r *RechargeAPI) CheckPayment(req *WebhookRequest) *WebhookResponse {
	states := req.State
	_, ok := states["order_id"]
	if !ok {
		res := NewWebhookResponse("Order_Failure")
		res.AddVariable("failureInfo", "Order Failed due to some internal error")
		return &res
	}
	res := NewWebhookResponse("Order_Success")
	res.AddVariable("successInfo", "Payment is successful")
	return &res
}

func CallBack(next_block string) *WebhookResponse {
	res := NewWebhookResponse(next_block)
	return &res
}
