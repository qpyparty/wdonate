package wdonate

// Request is a common structure for API requests containing authorization fields.
type Request struct {
	Token string `json:"token"`
	BotID int    `json:"botId"`
}

// GetLinkRequest represents a request to the getLink API method.
type GetLinkRequest struct {
	Request
	Sum               float64 `json:"sum,omitempty"`
	UserID            int     `json:"userId,omitempty"`
	Payload           int     `json:"payload,omitempty"`
	PriorityPayMethod string  `json:"priority_pay_method,omitempty"`
}

// GetPaymentsRequest represents a request to the getPayments API method.
type GetPaymentsRequest struct {
	Request
	Count int `json:"count"`
}

// GetBalanceRequest represents a request to the getBalance API method.
type GetBalanceRequest struct {
	Request
}

// GetCallbackRequest represents a request to the getCallback API method.
type GetCallbackRequest struct {
	Request
}

// SetCallbackRequest represents a request to the setCallback API method.
type SetCallbackRequest struct {
	Request
	CallbackURL string `json:"callbackUrl"`
}

// DelCallbackRequest represents a request to the delCallback API method.
type DelCallbackRequest struct {
	Request
}

// GetLinkResponse represents a response from the getLink API method.
type GetLinkResponse struct {
	Status   string `json:"status"`
	Response struct {
		Link string `json:"link"`
	} `json:"response"`
}

// GetPaymentsResponse represents a response from the getPayments API method.
type GetPaymentsResponse struct {
	Status   string `json:"status"`
	Response []struct {
		ID      int     `json:"id"`
		UserID  int     `json:"userId"`
		Type    string  `json:"type"`
		Sum     float64 `json:"sum"`
		TxnID   string  `json:"txnId"`
		Time    int     `json:"time"`
		Payload int     `json:"payload"`
	} `json:"response"`
}

// GetBalanceResponse represents a response from the getBalance API method.
type GetBalanceResponse struct {
	Status   string `json:"status"`
	Response struct {
		Balance float64 `json:"balance"`
	} `json:"response"`
}

// GetCallbackResponse represents a response from the getCallback API method.
type GetCallbackResponse struct {
	Status   string `json:"status"`
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}

// SetCallbackResponse represents a response from the setCallback API method.
type SetCallbackResponse struct {
	Status   string `json:"status"`
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}

// DelCallbackResponse represents a response from the delCallback API method.
type DelCallbackResponse struct {
	Status   string `json:"status"`
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}
