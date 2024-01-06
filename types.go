package wdonate

// Request is a common request structure with Token and BotID fields.
type Request struct {
	Token string `json:"token"`
	BotID int    `json:"botId"`
}

// Response is a common response structure with a Status field.
type Response struct {
	Status string `json:"status"` // Request status, success or error
}

// GetLinkRequest represents the request structure for the getLink API method.
type GetLinkRequest struct {
	Request
	Sum               float64 `json:"sum,omitempty"`                 // Default amount for the payment
	UserID            int     `json:"userId,omitempty"`              // User ID
	Payload           int     `json:"payload,omitempty"`             // Arbitrary number
	PriorityPayMethod string  `json:"priority_pay_method,omitempty"` // Preferred payment method
}

// GetLinkResponse represents the response structure for the getLink API method.
type GetLinkResponse struct {
	Response
	Link string `json:"link"` // Payment page link
}

// GetPaymentsRequest represents the request structure for the getPayments API method.
type GetPaymentsRequest struct {
	Request
	Count int `json:"count"` // Limit of records
}

// PaymentInfo represents the payment information in the getPayments response.
type PaymentInfo struct {
	ID      int     `json:"id"`      // Payment number in the system
	UserID  int     `json:"userId"`  // Payer's VK ID (if provided)
	Type    string  `json:"type"`    // Payment method
	Sum     float64 `json:"sum"`     // Payment amount in rubles
	TxnID   string  `json:"txnId"`   // Transaction ID in the payment system
	Time    int64   `json:"time"`    // Payment date and time in Unixtime
	Payload int     `json:"payload"` // Arbitrary number
}

// GetPaymentsResponse represents the response structure for the getPayments API method.
type GetPaymentsResponse struct {
	Response
	Payments []PaymentInfo `json:"response"`
}

// GetBalanceRequest represents the request structure for the getBalance API method.
type GetBalanceRequest Request

// GetBalanceResponse represents the response structure for the getBalance API method.
type GetBalanceResponse struct {
	Response
	Balance float64 `json:"balance"` // Current balance
}

// GetCallbackRequest represents the request structure for the getCallback API method.
type GetCallbackRequest Request

// GetCallbackResponse represents the response structure for the getCallback API method.
type GetCallbackResponse struct {
	Response
	URL string `json:"url"` // Current Callback URL
}

// SetCallbackRequest represents the request structure for the setCallback API method.
type SetCallbackRequest struct {
	Request
	CallbackURL string `json:"callbackUrl"` // Callback URL
}

// SetCallbackResponse represents the response structure for the setCallback API method.
type SetCallbackResponse struct {
	Response
	URL string `json:"url"` // New Callback URL
}

// DelCallbackRequest represents the request structure for the delCallback API method.
type DelCallbackRequest Request

// DelCallbackResponse represents the response structure for the delCallback API method.
type DelCallbackResponse struct {
	Response
	URL string `json:"url"` // Current Callback URL after deletion
}
