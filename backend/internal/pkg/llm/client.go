package llm

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

type ChatRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ChatResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}

func NewClient(baseURL, apiKey string) *Client {
    return &Client{
        BaseURL: baseURL,
        APIKey:  apiKey,
        HTTP:    &http.Client{Timeout: 60 * time.Second},
    }
}

func (c *Client) GenerateText(prompt string) (string, error) {
    payload := ChatRequest{
        Model: "llm-generic",
        Messages: []Message{{Role: "user", Content: prompt}},
    }
    body, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/chat/completions", bytes.NewReader(body))
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
        return "", fmt.Errorf("llm request failed: %s", resp.Status)
    }

    var result ChatResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    if len(result.Choices) == 0 {
        return "", fmt.Errorf("llm response is empty")
    }
    return result.Choices[0].Message.Content, nil
}

func (c *Client) GenerateJSON(prompt string, out any) error {
    content, err := c.GenerateText(prompt)
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(content), out)
}
