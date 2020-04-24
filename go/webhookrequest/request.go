package webhookrequest

//WebhookRequest is a sample request from wario
type WebhookRequest struct {
	CallbackURL      string `json:"callback_url"`
	CurrentBlock     string `json:"current_block"`
	Expiry           int    `json:"expiry"`
	ID               string `json:"id"`
	RoomCode         int    `json:"room_code"`
	RoomCustomFields struct {
	} `json:"room_custom_fields"`
	RoomType  string                         `json:"room_type"`
	Source    string                         `json:"source"`
	State     map[string]string              `json:"state,omitempty"`
	UserID    string                         `json:"user_id"`
	Variables map[string](map[string]string) `json:"variables"`
	Visitor   struct {
		Avatar string `json:"avatar"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Phone  string `json:"phone"`
	} `json:"visitor"`
	VisitorCustomFields map[string]string `json:"visitor_custom_fields,omitempty"`
}
