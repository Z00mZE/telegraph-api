package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	url    string
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: new(http.Client),
	}
}

type rawResp struct {
	Ok     bool            `json:"ok"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func (c *Client) Send(ctx context.Context, url string, request, response any) error {
	payload, payloadError := json.Marshal(request)

	if payloadError != nil {
		return payloadError
	}

	req, reqError := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	if reqError != nil {
		return reqError
	}

	resp, respError := c.client.Do(req)
	if respError != nil {
		return respError
	}
	defer resp.Body.Close()

	body, bodyError := io.ReadAll(resp.Body)
	if bodyError != nil {
		return bodyError
	}

	raw := new(rawResp)
	if parseError := json.Unmarshal(body, raw); parseError != nil {
		return parseError
	}

	if !raw.Ok {
		if raw.Error != "" {
			return errors.New(raw.Error)
		}
		return UndefinedError
	}

	return json.Unmarshal(raw.Result, response)
}
