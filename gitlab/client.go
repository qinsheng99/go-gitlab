package gitlab

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	c *http.Client
}

type Token struct {
	token string
}

type WithToken func(t *Token)

func Tokenstr(token string) WithToken {
	return func(t *Token) {
		t.token = token
	}
}

func NewToken(options ...WithToken) *Token {
	var t = &Token{}
	for _, option := range options {
		option(t)
	}
	return t
}

func (t *Token) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(t.token) == 0 {
		return nil, errors.New("token is nil")
	}

	r := t.newReq(req)
	r.Header.Set("Authorization", "Bearer"+" "+t.token)

	return http.DefaultTransport.RoundTrip(r)
}

func (t *Token) newReq(r *http.Request) *http.Request {
	r1 := new(http.Request)
	*r1 = *r
	r1.Header = make(http.Header, len(r.Header))
	for s, h := range r.Header {
		r1.Header[s] = append([]string(nil), h...)
	}

	return r1

}

const base = "https://gitlab.com/api/v4/"
const max = 3

func NewClient(c *http.Client) Inter {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{c: c}
}

func (c *Client) RetCli() *http.Client {
	return c.c
}

func (c *Client) request(ctx context.Context, method, u string, heads map[string]string, body io.Reader, values url.Values) (*http.Response, error) {
	var (
		err  error
		newh *http.Request
		do   *http.Response
		path *url.URL
	)

	path, err = url.Parse(u)

	if len(values) > 0 {
		q := path.Query()

		for s, value := range values {
			for _, v := range value {
				q.Add(s, v)
			}
		}
		path.RawQuery = q.Encode()
	}

	if body != nil {
		newh, err = http.NewRequestWithContext(ctx, method, path.String(), body)
	} else {
		newh, err = http.NewRequestWithContext(ctx, method, path.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	if len(heads) > 0 {
		headers := http.Header{}
		for h, v := range heads {
			headers.Set(h, v)
		}

		newh.Header = headers
	}

	for i := 0; i < max; i++ {
		do, err = c.c.Do(newh)
		if err != nil {
			if i+1 >= max {
				return nil, errors.New("too many request failed")
			}
			time.Sleep(3 * time.Second)
			continue
		}
		if err == nil {
			break
		}
	}

	return do, nil

}
