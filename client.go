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
	Token string
	BotID int
}

// NewClient creates a new instance of the Wdonate API client.
func NewClient(token string, botId int) *Client {
	return &Client{
		Token: token,
		BotID: botId,
	}
}

// doRequest performs a POST request to the Wdonate API with the given endpoint and request body.
func (c *Client) doRequest(endpoint string, requestBody interface{}) ([]byte, error) {
	url := fmt.Sprintf(BASE_API_URL, endpoint)
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request to %s failed with status code: %d", endpoint, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// GetLink sends a request to the getLink API method and returns the response.
func (c *Client) GetLink(req GetLinkRequest) (*GetLinkResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("getLink", req)
	if err != nil {
		return nil, err
	}

	var resp GetLinkResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

// GetPayments sends a request to the getPayments API method and returns the response.
func (c *Client) GetPayments(req GetPaymentsRequest) (*GetPaymentsResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("getPayments", req)
	if err != nil {
		return nil, err
	}

	var resp GetPaymentsResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

// GetBalance sends a request to the getBalance API method and returns the response.
func (c *Client) GetBalance(req GetBalanceRequest) (*GetBalanceResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("getBalance", req)
	if err != nil {
		return nil, err
	}

	var resp GetBalanceResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

// GetCallback sends a request to the getCallback API method and returns the response.
func (c *Client) GetCallback(req GetCallbackRequest) (*GetCallbackResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("getCallback", req)
	if err != nil {
		return nil, err
	}

	var resp GetCallbackResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

// SetCallback sends a request to the setCallback API method and returns the response.
func (c *Client) SetCallback(req SetCallbackRequest) (*SetCallbackResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("setCallback", req)
	if err != nil {
		return nil, err
	}

	var resp SetCallbackResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}

// DelCallback sends a request to the delCallback API method and returns the response.
func (c *Client) DelCallback(req DelCallbackRequest) (*DelCallbackResponse, error) {
	req.Token = c.Token
	req.BotID = c.BotID

	body, err := c.doRequest("delCallback", req)
	if err != nil {
		return nil, err
	}

	var resp DelCallbackResponse
	err = json.Unmarshal(body, &resp)
	return &resp, err
}
