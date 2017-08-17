package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const defaultTimeout = 10 * time.Second

// 共享的http transport
var transport = &http.Transport{
	MaxIdleConnsPerHost: 8,
}

// Client http客户端结构
type Client struct {
	method      string
	url         string
	header      map[string][]string
	body        []byte
	contentType string
	timeout     time.Duration
}

// Response http客户端返回响应结构
type Response struct {
	StatusCode int
	Body       []byte
}

// Option 客户端调用选项
type Option func(*Client)

// ContentType 设置客户端content_type
func ContentType(contentType string) Option {
	return func(c *Client) {
		c.contentType = contentType
	}
}

// ContentJSON 设置客户端content_type为application/json
func ContentJSON() Option {
	return ContentType("application/json")
}

// Timeout 设置客户端超时时间，包括连接时间，重定向时间，读响应Body的时间
func Timeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// Header 设置客户端请求头部
func Header(key, value string) Option {
	return func(c *Client) {
		if c.header == nil {
			c.header = make(map[string][]string)
		}
		c.header[key] = append(c.header[key], value)
	}
}

// Head Head方法调用
func Head(url string, options ...Option) (*Response, error) {
	c := &Client{method: "HEAD", url: url}
	return c.do(options...)
}

// Get Get方法调用
func Get(url string, options ...Option) (*Response, error) {
	c := &Client{method: "GET", url: url}
	return c.do(options...)
}

// Post Post方法调用
func Post(url string, body []byte, options ...Option) (*Response, error) {
	c := &Client{method: "POST", url: url, body: body}
	return c.do(options...)
}

// Put Put方法调用
func Put(url string, body []byte, options ...Option) (*Response, error) {
	c := &Client{method: "PUT", url: url, body: body}
	return c.do(options...)
}

func (c *Client) do(options ...Option) (*Response, error) {
	for _, option := range options {
		option(c)
	}

	req, err := http.NewRequest(c.method, c.url, bytes.NewReader(c.body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if c.contentType != "" {
		req.Header.Add("Content-Type", c.contentType)
	}

	if c.timeout == 0 {
		c.timeout = defaultTimeout
	}

	if c.header != nil {
		for key, values := range c.header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	client := http.Client{
		Transport: transport,
		Timeout:   c.timeout,
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rsp.Body.Close()

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Response{StatusCode: rsp.StatusCode, Body: rspBody}, nil
}
