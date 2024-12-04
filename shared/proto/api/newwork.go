package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/url"
)

type BaseApi struct {
	Ctx context.Context
	Url string
}

func New(ctx context.Context) *BaseApi {
	return &BaseApi{Ctx: ctx}
}

// Post 发送 POST 请求
func (api *BaseApi) Post(url string, param interface{}, resp interface{}) error {
	Url := api.Url + url
	jsonData, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}
	if err := json.Unmarshal(body, resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}
