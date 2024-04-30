package http_client

import (
	"bytes"
	"errors"
	"net/http"
	netUrl "net/url"
	"sync"
	"time"
)

type FnHeaderConfig func(header *http.Header)

type Option func(*HttpClient)

func WithTimeout(timeout time.Duration) Option {
	return func(hc *HttpClient) {
		hc.timeout = timeout
	}
}

func WithRedirectMaxTimes(maxTimes uint8) Option {
	return func(hc *HttpClient) {
		hc.redirectMaxTimes = maxTimes
	}
}

func WithRetryMaxTimes(maxTimes uint8) Option {
	return func(hc *HttpClient) {
		hc.retryMaxTimes = maxTimes
	}
}
func WithCookie(cookie string) Option {
	return func(hc *HttpClient) {
		hc.cookie = cookie
	}
}

type HttpClient struct {
	client           *http.Client
	cookie           string
	m                sync.Mutex
	timeout          time.Duration
	redirectMaxTimes uint8
	retryMaxTimes    uint8
}

func NewHttpClient(options ...Option) *HttpClient {
	hc := &HttpClient{
		cookie:           "",
		client:           http.DefaultClient,
		timeout:          30 * time.Second,
		redirectMaxTimes: 5,
		retryMaxTimes:    3,
	}

	for _, option := range options {
		option(hc)
	}

	return hc
}

type HttpType string

const (
	POST = "POST"
	GET  = "GET"
)

func (hc *HttpClient) Do(httpType HttpType, url string, fnHeaderConfig FnHeaderConfig, queryMap map[string]string, body []byte) (*http.Response, error) {
	if queryMap != nil {
		parse, _ := netUrl.Parse(url)
		query := parse.Query()
		for key, value := range queryMap {
			query.Add(key, value)
		}
		parse.RawQuery = query.Encode()
		url = parse.String()
	}
	req, err := http.NewRequest(string(httpType), url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
	req.Header.Add("Referer", "https://www.bilibili.com/")

	if fnHeaderConfig != nil {
		fnHeaderConfig(&req.Header)
	}
	if hc.cookie != "" {
		req.Header.Set("cookie", hc.cookie)
	}
	return hc.doRequest(req)
}
func (hc *HttpClient) doRequest(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for i := uint8(0); i < hc.retryMaxTimes; i++ {
		resp, err = hc.client.Do(req)
		if err != nil {
			if i == hc.retryMaxTimes-1 {
				return nil, err
			}
			continue
		}
		break
	}
	if resp == nil {
		return nil, errors.New("empty response")
	}
	return resp, nil
}
