package onwater

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const apiEndpoint = "https://api.onwater.io/api/v1/results/"

// Client struct
type Client struct {
	apiKey string
}

// New client initalizes and returns new Client object that you can use to make API calls.
// If apiKey is empty, it tries to load key from $ONWATER_API_KEY environment variable.
func New(apiKey string) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("ONWATER_API_KEY")
	}
	return &Client{apiKey}
}

// OnWater makes an API call and returns true if lat/lng is on water.
func (c *Client) OnWater(ctx context.Context, lat float64, lng float64) (bool, error) {
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

// OnLand returns true if lat/lng is on land
func (c *Client) OnLand(ctx context.Context, lat float64, lng float64) (bool, error) {
	ok, err := c.OnWater(ctx, lat, lng)
	if err != nil {
		return ok, err
	}
	return !ok, err
}
