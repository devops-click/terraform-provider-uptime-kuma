// client.go
package main

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"errors"
)

type Client struct {
	URL      string
	Username string
	Password string
	APIKey   string
	Token    string
}

func (c *Client) Authenticate() error {
	if c.APIKey != "" {
		c.Token = c.APIKey
		return nil
	}

	data := map[string]string{
		"username": c.Username,
		"password": c.Password,
	}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(fmt.Sprintf("%s/api/login", c.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to authenticate")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	token, ok := result["token"].(string)
	if !ok {
		return errors.New("failed to get token from response")
	}
	c.Token = token
	return nil
}
