package econet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	baseURL                = "https://rheem.clearblade.com/api/v/1"
	clearBladeSystemKey    = "e2e699cb0bb0bbb88fc8858cb5a401"
	clearBladeSystemSecret = "E2E699CB0BE6C6FADDB1B0BC9A20"
)

type Client struct {
	UserID    string `json:"user_id"`
	UserToken string `json:"user_token"`
	HTTP      *http.Client
	Options   struct {
		CONNECTIVITY struct {
			LastConnected    int64 `json:"last_connected"`
			LastDisconnected int64 `json:"last_disconnected"`
		} `json:"@CONNECTIVITY"`
		ID                        string      `json:"_id"`
		AccountID                 string      `json:"account_id"`
		Address                   string      `json:"address"`
		AllowEmailNotifications   bool        `json:"allow_email_notifications"`
		AllowProductAlertEmails   bool        `json:"allow_product_alert_emails"`
		AllowProductAlertTextMsg  bool        `json:"allow_product_alert_text_msg"`
		AllowPushNotifications    bool        `json:"allow_push_notifications"`
		AllowSpecialOffersEmails  bool        `json:"allow_special_offers_emails"`
		AllowSpecialOffersTextMsg bool        `json:"allow_special_offers_text_msg"`
		AllowTextNotifications    bool        `json:"allow_text_notifications"`
		CbServiceAccount          bool        `json:"cb_service_account"`
		CbTTLOverride             int         `json:"cb_ttl_override"`
		City                      string      `json:"city"`
		Connected                 bool        `json:"connected"`
		Email                     string      `json:"email"`
		EmailValidated            bool        `json:"email_validated"`
		FirstName                 string      `json:"first_name"`
		IsPhoneVerified           bool        `json:"is_phone_verified"`
		LastName                  string      `json:"last_name"`
		Phone                     interface{} `json:"phone"`
		PhoneNumber               string      `json:"phone_number"`
		PhoneValidated            bool        `json:"phone_validated"`
		PostalCode                string      `json:"postal_code"`
		ReceiveMarketingMessages  bool        `json:"receive_marketing_messages"`
		ReportState               bool        `json:"report_state"`
		Role                      int         `json:"role"`
		ShareStatus               int         `json:"share_status"`
		State                     string      `json:"state"`
		Success                   bool        `json:"success"`
		TemperatureDisplayUnit    string      `json:"temperature_display_unit"`
		TwoFactorEnabled          bool        `json:"two_factor_enabled"`
		TwoFactorMethod           string      `json:"two_factor_method"`
		UserID                    string      `json:"user_id"`
	} `json:"options"`
}

func New(username string, password string) (*Client, error) {
	client := &Client{
		HTTP: &http.Client{},
	}
	return client.authorize(username, password)
}

func (c *Client) addHeaders(req *http.Request) {
	headers := map[string]string{
		"ClearBlade-SystemKey":    clearBladeSystemKey,
		"ClearBlade-SystemSecret": clearBladeSystemSecret,
		"Content-Type":            "application/json; charset=UTF-8",
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if c.UserToken != "" {
		req.Header.Set("ClearBlade-UserToken", c.UserToken)
	}
}

func (c *Client) authorize(username string, password string) (*Client, error) {
	url := baseURL + "/user/auth"

	requestBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    username,
		Password: password,
	}
	data, _ := json.Marshal(requestBody)

	body, err := c.doPost(data, url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) doPost(body []byte, url string) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	body, err := c.processRequest(req)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (c *Client) processRequest(req *http.Request) ([]byte, error) {
	c.addHeaders(req)

	// send request
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) getLocations() (*Locations, error) {
	url := baseURL + "/code/" + clearBladeSystemKey + "/getLocation"

	body, err := c.doPost([]byte{}, url)
	if err != nil {
		return nil, err
	}

	locations := &Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
