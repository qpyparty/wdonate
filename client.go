package wdonate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client represents a client for the Wdonate API.
type Client struct {
	Token      string
	BotID      int
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new instance of Client.
func NewClient(token string, botID int) *Client {
	return &Client{
		Token:      token,
		BotID:      botID,
		BaseURL:    BASE_API_URL,
		HTTPClient: &http.Client{},
	}
}

// doRequest performs a HTTP request and decodes the response.
func (c *Client) doRequest(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

// GetLink sends a request to the getLink API method and decodes the response.
func (c *Client) GetLink(request *GetLinkRequest) (*GetLinkResponse, error) {
	if request == nil {
		request = &GetLinkRequest{}
	}
	if request.Token == "" {
		request.Token = c.Token
	}
	if request.BotID == 0 {
		request.BotID = c.BotID
	}

	url := fmt.Sprintf(c.BaseURL, "getLink")
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response GetLinkResponse
	err = c.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPayments sends a request to the getPayments API method and decodes the response.
func (c *Client) GetPayments(request *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	if request == nil {
		request = &GetPaymentsRequest{}
	}
	if request.Token == "" {
		request.Token = c.Token
	}
	if request.BotID == 0 {
		request.BotID = c.BotID
	}

	url := fmt.Sprintf(c.BaseURL, "getPayments")
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response GetPaymentsResponse
	err = c.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalance sends a request to the getBalance API method and decodes the response.
func (c *Client) GetBalance(request *GetBalanceRequest) (*GetBalanceResponse, error) {
	if request == nil {
		request = &GetBalanceRequest{}
	}
	if request.Token == "" {
		request.Token = c.Token
	}
	if request.BotID == 0 {
		request.BotID = c.BotID
	}

	url := fmt.Sprintf(c.BaseURL, "getBalance")
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response GetBalanceResponse
	err = c.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCallback sends a request to the getCallback API method and decodes the response.
func (c *Client) GetCallback(request *GetCallbackRequest) (*GetCallbackResponse, error) {
	if request == nil {
		request = &GetCallbackRequest{}
	}
	if request.Token == "" {
		request.Token = c.Token
	}
	if request.BotID == 0 {
		request.BotID = c.BotID
	}

	url := fmt.Sprintf(c.BaseURL, "getCallback")
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response GetCallbackResponse
	err = c.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SetCallback sends a request to the setCallback API method and decodes the response.
func (c *Client) SetCallback(request *SetCallbackRequest) (*SetCallbackResponse, error) {
	if request == nil {
		request = &SetCallbackRequest{}
	}
	if request.Token == "" {
		request.Token = c.Token
	}
	if request.BotID == 0 {
		request.BotID = c.BotID
	}

	url := fmt.Sprintf(c.BaseURL, "setCallback")
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response SetCallbackResponse
	err = c.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
