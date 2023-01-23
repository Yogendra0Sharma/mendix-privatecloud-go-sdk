package mendix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	defaultBaseURL  = "https://privatecloud.mendixcloud.com/"
	apiVersionPath  = "api/v3/"
	DefaultTimeout  = 30 * time.Second
	DefaultMaxConns = 50
)

// Client represents the base for connecting to Mendix Private Cloud
type Client struct {
	// Base URL for API requests. Defaults to the mendix private cloud api
	baseURL *url.URL
	// HTTP client used to communicate with the API.
	client *http.Client
	// Token (MxToken) used to make authenticated API calls.
	token string
	// CLI API Service
	Cli *CLIService
	// Cluster API Service
	Cluster *ClusterService
	// NameSpace API Service
	Namespace *NameSpaceService
	// Environment API Service
	Environment *EnvironmentService
}

func NewClient(token string) *Client {
	c := &Client{}
	c.client = &http.Client{Timeout: DefaultTimeout}
	if token == "" {
		c.token = os.Getenv("MX_TOKEN")
	} else {
		c.token = token
	}
	// Set the default base URL.
	c.setBaseURL(defaultBaseURL)
	// setup API services
	c.Cli = &CLIService{client: c}
	c.Cluster = &ClusterService{client: c}
	c.Namespace = &NameSpaceService{client: c}
	c.Environment = &EnvironmentService{client: c}
	return c
}

// BaseURL return a copy of the baseURL.
func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

func (c *Client) setAuthHeader(req *http.Request) error {
	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("MxToken %s", c.token))
		return nil
	}
	return errors.New("missing Authorization Header")
}

func (c *Client) setCommonHeaders(req *http.Request) error {
	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")
	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}
	return errors.New("missing Authorization Header")
}

func (c *Client) NewRequest(method, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	var reqBody io.Reader = nil

	if method == http.MethodPost || method == http.MethodPut {
		if opt != nil {
			marshalled, err := json.Marshal(opt)
			reqBody = bytes.NewReader(marshalled)
			if err != nil {
				return nil, err
			}
		}
	}
	req, err := http.NewRequest(method, u.String(), reqBody)
	c.setAuthHeader(req)
	c.setCommonHeaders(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(
	req *http.Request,
	result interface{}) (*http.Response, error) {
	res, err := c.sendRequest(req, result)
	return res, err
}

func (c *Client) sendRequest(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}
