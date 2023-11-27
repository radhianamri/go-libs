package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/radhianamri/go-libs/cache"
	"github.com/radhianamri/go-libs/errors"
)

type Client struct {
	HttpClient *http.Client
	Cache      cache.Cache
}

type RequestOpt struct {
	Body     []byte
	Header   map[string]string
	Query    url.Values
	CacheOpt CacheOpt
}

type CacheOpt struct {
	Enabled bool
	Key     string
	TTL     time.Duration
}

type Response struct {
	Body       []byte
	StatusCode int
}

func (c *Client) SendRequest(ctx context.Context, method methodName, url string, requestOpt *RequestOpt) (*Response, error) {
	if requestOpt.CacheOpt.Enabled {
		cacheBody, err := c.Cache.Get(ctx, requestOpt.CacheOpt.Key)
		if err == nil {
			return &Response{
				Body:       cacheBody,
				StatusCode: http.StatusOK,
			}, nil
		}
	}

	req, err := http.NewRequestWithContext(ctx, string(method), url, bytes.NewBuffer(requestOpt.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %s", err)
	}

	if requestOpt != nil {
		for k, v := range requestOpt.Header {
			req.Header.Set(k, v)
		}
		req.URL.RawQuery = requestOpt.Query.Encode()
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %s", err)
	}

	go c.Cache.Set(ctx, requestOpt.CacheOpt.Key, respBody, requestOpt.CacheOpt.TTL)

	return &Response{
		Body:       respBody,
		StatusCode: resp.StatusCode,
	}, nil
}

type RetryOpt struct {
	Count  int
	Delay  time.Duration
	Errors []error
}

func (c *Client) SendRequestWithRetry(ctx context.Context, method methodName, url string, requestOpt *RequestOpt, retryOpt RetryOpt) (resp *Response, err error) {
	for i := 1; i <= retryOpt.Count; i++ {
		resp, err := c.SendRequest(ctx, method, url, requestOpt)
		if err == nil || !errors.Contains(retryOpt.Errors, err) {
			return resp, err
		}

		if i == retryOpt.Count {
			break
		}
		time.Sleep(retryOpt.Delay)
	}

	return nil, fmt.Errorf("failed to execute request after %d retries: %w", retryOpt.Count, err)
}
