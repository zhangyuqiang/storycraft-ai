package seedance

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Client struct {
    BaseURL string
    APIKey  string
    HTTP    *http.Client
}

type ImageRequest struct {
    Prompt string `json:"prompt"`
    Model  string `json:"model,omitempty"`
}

type ImageResponse struct {
    Data []struct {
        URL string `json:"url"`
    } `json:"data"`
}

type VideoRequest struct {
    Prompt string `json:"prompt"`
    Model  string `json:"model,omitempty"`
    Width  int    `json:"width,omitempty"`
    Height int    `json:"height,omitempty"`
}

type VideoResponse struct {
    Data []struct {
        URL string `json:"url"`
    } `json:"data"`
}

func NewClient(baseURL, apiKey string) *Client {
    return &Client{
        BaseURL: baseURL,
        APIKey:  apiKey,
        HTTP:    &http.Client{Timeout: 90 * time.Second},
    }
}

func (c *Client) GenerateImage(prompt string) (string, error) {
    payload := ImageRequest{Prompt: prompt, Model: "seedance-image"}
    body, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/images/generations", bytes.NewReader(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    if c.APIKey != "" {
        req.Header.Set("Authorization", "Bearer "+c.APIKey)
    }

    resp, err := c.HTTP.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return "", fmt.Errorf("seedance image request failed: %s", resp.Status)
    }

    var result ImageResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    if len(result.Data) == 0 || result.Data[0].URL == "" {
        return "", fmt.Errorf("seedance image returned no URL")
    }
    return result.Data[0].URL, nil
}

func (c *Client) GenerateVideo(prompt string) (string, error) {
    payload := VideoRequest{Prompt: prompt, Model: "seedance-video"}
    body, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/videos/generations", bytes.NewReader(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    if c.APIKey != "" {
        req.Header.Set("Authorization", "Bearer "+c.APIKey)
    }

    resp, err := c.HTTP.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return "", fmt.Errorf("seedance video request failed: %s", resp.Status)
    }

    var result VideoResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    if len(result.Data) == 0 || result.Data[0].URL == "" {
        return "", fmt.Errorf("seedance video returned no URL")
    }
    return result.Data[0].URL, nil
}
