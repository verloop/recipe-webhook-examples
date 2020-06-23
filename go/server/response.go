package server

import "strings"

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
	Title    string   `json:"title"`
	Action   Actions  `json:"action"`
	Triggers []string `json:"triggers,omitempty"`
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
		Title:    title,
		Action:   Action,
		Triggers: make([]string, 0),
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

//AddTriggers adds a list of keywords which would trigger the action
func (this *Operators) AddTriggers(t ...string) {
	for _, k := range t {
		this.Triggers = append(this.Triggers, k)
	}
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
