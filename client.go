package gopay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	goId            int
	tokenExpiration time.Time
	token           string
	clientID        string
	clientSecret    string
	URL             string
}

type requestBodyJSON struct {
	body interface{}
}

func (rbj requestBodyJSON) toBytes() ([]byte, error) {
	return json.Marshal(rbj.body)
}

func (rbf requestBodyJSON) contentType() string {
	return "application/json"
}

type requestBodyForm struct {
	body url.Values
}

func (rbf requestBodyForm) toBytes() ([]byte, error) {
	return []byte(rbf.body.Encode()), nil
}

func (rbf requestBodyForm) contentType() string {
	return "application/x-www-form-urlencoded"
}

type RequestBody interface {
	toBytes() ([]byte, error)
	contentType() string
}

func (c *Client) doRequest(path string, method string, body RequestBody) (*http.Response, error) {
	err := c.RefreshToken()
	if err != nil {
		return nil, err
	}

	contentType := "application/x-www-form-urlencoded"
	var finalBody *bytes.Buffer = bytes.NewBuffer([]byte{})
	if body != nil {
		contentType = body.contentType()
		_bytes, err := body.toBytes()
		if err != nil {
			return nil, err
		}

		finalBody = bytes.NewBuffer(_bytes)
	}

	req, err := http.NewRequest(method, c.URL+path, finalBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			body = []byte("failed to get body")
		}

		return nil, fmt.Errorf("Reponse status is not 200 (status: %s, statusCode: %d) %s", res.Status, res.StatusCode, string(body))
	}

	return res, nil
}

type RefreshTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) RefreshToken() error {
	// Token ještě nepropadl
	if c.tokenExpiration.After(time.Now()) {
		return nil
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "payment-all")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth2/token", c.URL), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	responseData := RefreshTokenResponse{}

	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	c.tokenExpiration = time.Now().Add(time.Duration(responseData.ExpiresIn) * time.Second)
	c.token = responseData.AccessToken

	return nil
}

func NewClient(id string, secret string, goId int, url string) Client {
	return Client{
		tokenExpiration: time.Now(),
		token:           "",
		clientSecret:    secret,
		goId:            goId,
		URL:             url,
		clientID:        id,
	}
}
