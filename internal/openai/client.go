package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/pkg/errors"
)

const (
	apiTokenHeader = "X-JFrog-Art-Api"
	answerPath     = "/answers"
)

type Client struct {
	url        *url.URL
	token      string
	httpClient *http.Client
}

func NewClient(u *url.URL, token string) (*Client, error) {
	return &Client{
		url:        u,
		token:      token,
		httpClient: cleanhttp.DefaultClient(),
	}, nil
}

type Answer struct {
	Answers []string
	Text    string
}

// CreateAnswer Answers the specified question using the provided documents and examples.
// https://beta.openai.com/docs/api-reference/answers/create
func (c *Client) CreateAnswer(ctx context.Context, q string) (Answer, error) {
	var payload struct {
		File     string `json:"file"`
		Question string `json:"question"`
	}

	if err := c.do(ctx, http.StatusOK, http.MethodGet, u, nil, &payload); err != nil {
		return Answer{}, err
	}
}

func (c *Client) do(ctx context.Context, wantStatus int, method string, u *url.URL, body io.Reader, payload interface{}) error {
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return errors.Wrap(err, "creating request failed")
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "create answer failed")
	}
	defer res.Body.Close()

	if want, got := wantStatus, res.StatusCode; want != got {
		return errors.Errorf("wanted status: %d, got %d", want, got)
	}

	return json.NewDecoder(res.Body).Decode(payload)
}
