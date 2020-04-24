package webhookresponse

import (
	"strings"

	"github.com/verloop/recipe-webhook-examples/go/webhookrequest"
)

//WebhookResponse has info about next block and quick replies
type WebhookResponse struct {
	NextBlock string                   `json:"next_block,omitempty"`
	State     map[string]string        `json:"state,omitempty"`
	Variable  map[string]string        `json:"variables,omitempty"`
	Exports   map[string][]interface{} `json:"exports,omitempty"`
}

//NewWebhookResponse creates a webhook response for given next block string
func NewWebhookResponse(next string) WebhookResponse {
	return WebhookResponse{
		NextBlock: next,
		State:     make(map[string]string),
		Variable:  make(map[string]string),
		Exports:   make(map[string][]interface{}),
	}
}

//AddExportList adds a list of entries to given key
func (this *WebhookResponse) AddExportList(key string, vslice ...interface{}) {
	for _, v := range vslice {
		this.Exports[key] = append(this.Exports[key], v)
	}
}

func (this *WebhookResponse) AddState(key string, v string) {
	this.State[key] = v
}

func (this *WebhookResponse) AddVariable(key string, v string) {
	this.Variable[key] = v
}

type Operators struct {
	Title  string  `json:"title"`
	Action Actions `json:"action"`
}

func checkValidOperator(name string) bool {
	for _, k := range operatorList {
		if strings.ToLower(name) == k {
			return true
		}
	}
	return false
}

var operatorList []string

func init() {
	operatorList = []string{"vodafone", "airtel"} //for now only 2
}

func NewOperators(title string, Action Actions) Operators {
	return Operators{
		Title:  title,
		Action: Action,
	}
}

//Actions represent next block to excute when clicked upon and also holds custom variables
type Actions struct {
	NextBlock string            `json:"next_block"`
	Variables map[string]string `json:"variables,omitempty"`
}

func NewActions(blockname string) Actions {
	return Actions{
		NextBlock: blockname,
		Variables: make(map[string]string),
	}
}

//AddVariable add any list of custom variables pertaining to the action
func (this *Actions) AddVariable(key, value string) {
	this.Variables[key] = value
}

//Plan hold different plans of operator, Buttons are slice of web_url or postback buttons
type Plan struct {
	Title    string        `json:"title"`
	SubTitle string        `json:"subtitle"`
	Buttons  []interface{} `json:"buttons"`
}

//AddButton adds button/list of type postback or web_url
func (this *Plan) AddButton(value ...interface{}) {
	for _, v := range value {
		this.Buttons = append(this.Buttons, v)
	}
}

//NewPlan returns are new plan
func NewPlan(t, s string, b []interface{}) Plan {
	return Plan{
		Title:    t,
		Buttons:  b,
		SubTitle: s,
	}
}

//Button is poskback button for now
type Button struct {
	Title  string  `json:"title"`
	Type   string  `json:"type"`
	Action Actions `json:"action"`
}

//NewButton returns a new button
func NewButton(title, typ string, a Actions) Button {
	return Button{
		Title:  title,
		Type:   typ,
		Action: a,
	}
}

//URL is of web_url type button
type URL struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

//NewButton returns a new URL
func NewURL(title, typ, url string) URL {
	return URL{
		Title: title,
		Type:  typ,
		Url:   url,
	}
}

//ConstructOperators returns list of all operators
func ConstructOperators() *WebhookResponse {
	res := NewWebhookResponse("Show_Operators")
	a1 := NewActions("Get_Plans_Wait")
	a1.AddVariable("operator", "Vodafone")
	a2 := NewActions("Get_Plans_Wait")
	a2.AddVariable("operator", "Airtel")
	res.AddExportList("OperatorList", NewOperators("Vodafone", a1))
	res.AddExportList("OperatorList", NewOperators("Airtel", a2))
	return &res
}

//ConstructPlansForOperator constructs a recharge plan
func ConstructPlansForOperator(req *webhookrequest.WebhookRequest) *WebhookResponse {

	// operatorKey, ok := req.Variables["operator"]
	// if !ok {
	// 	log.Println("operator key not present")
	// 	req := NewWebhookResponse("Invalid_Operator")
	// 	return &req
	// }

	// operator, ok := operatorKey["parsed_value"]

	// if !ok || !checkValidOperator(operator) {
	// 	log.Println("invalid operator")
	// 	req := NewWebhookResponse("Invalid_Operator")
	// 	return &req
	// }

	res := NewWebhookResponse("Show_Plans_Text")

	a := NewActions("Do_Recharge")
	a.AddVariable("amount", "100")
	button := NewButton("Select", "postback", a)
	url := NewURL("Know More", "web_url", "https://verloop.io")

	plan1 := NewPlan("Data 1 GB (28 days)", "Rs. 100", nil)
	// plan1.AddButton(button)
	// plan1.AddButton(url)
	plan1.AddButton(button, url)

	a = NewActions("Do_Recharge")
	a.AddVariable("amount", "150")
	button = NewButton("Select", "postback", a)

	plan2 := NewPlan("Full Talk time (84 days)", "Rs. 150", nil)
	// plan2.AddButton(button)
	// plan2.AddButton(url)
	plan2.AddButton(button, url)

	//res.AddExportList("PlanList", plan1)
	//res.AddExportList("PlanList", plan2)
	res.AddExportList("PlanList", plan1, plan2)
	return &res
}

//ConstructPaymentLink returns a dummy payment link
func ConstructPaymentLink(req *webhookrequest.WebhookRequest) *WebhookResponse {
	// amountkey, ok := req.Variables["amount"]
	// if !ok {
	// 	res := NewWebhookResponse("Invalid_Options")
	// 	return &res
	// }

	// amount, ok := amountkey["value"]
	// if !ok {
	// 	res := NewWebhookResponse("Invalid_Options")
	// 	return &res
	// }
	amount := "100"
	res := NewWebhookResponse("")
	res.AddState("order_id", "NXPAOMDAJDAY")

	button := NewButton("Payment Done", "postback", NewActions("Verify_Payment"))
	url := NewURL("Pay Rs : "+amount+" INR", "web_url", "https://verloop.io/pricing.html")

	//res.AddExportList("PaymentOptions", button)
	//res.AddExportList("PaymentOptions", url)
	res.AddExportList("PaymentOptions", button, url)
	return &res
}

//CheckPayment checks for a order_id in the request state
//If order_id is present it returns Order_Success to be executed next else returns Order_Failure block to be executed next
func CheckPayment(req *webhookrequest.WebhookRequest) *WebhookResponse {
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
