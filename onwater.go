package onwater

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const apiEndpoint = "https://api.onwater.io/api/v1/results/"

// Client struct
type Client struct {
	apiKey string
}

// New client
func New(apiKey string) *Client {
	return &Client{apiKey}
}

// OnWater returns true if lat/lng are on water
func (c *Client) OnWater(ctx context.Context, lat float32, lng float32) (bool, error) {
	url := fmt.Sprintf("%s/%v,%v", apiEndpoint, lat, lng)
	if c.apiKey != "" {
		url += "?access_token=" + c.apiKey
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}
	req.Header.Add("User-Agent", "github.com/andy/onwater-1.0.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, errors.New(res.Status)
	}

	var reply struct {
		Lat   float32 `json:"lat"`
		Lon   float32 `json:"lon"`
		Water bool    `json:"water"`
	}

	if err := json.NewDecoder(res.Body).Decode(&reply); err != nil {
		return false, err
	}

	return reply.Water, nil
}

// OnLand returns true if lat/lng are on land
func (c *Client) OnLand(ctx context.Context, lat float32, lng float32) (bool, error) {
	ok, err := c.OnWater(ctx, lat, lng)
	if err != nil {
		return ok, err
	}
	return !ok, err
}
