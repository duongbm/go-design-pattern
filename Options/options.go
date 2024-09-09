package OptionsPattern

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Requester struct {
	BaseUrl  string
	Timeout  time.Duration
	Data     map[string]interface{}
	Headers  map[string]interface{}
	Response interface{}
}

type RequestOption func(*Requester)

func WithTimeout(seconds int) RequestOption {
	return func(r *Requester) {
		r.Timeout = time.Duration(seconds) * time.Second
	}
}

func WithHeader(key, value string) RequestOption {
	return func(r *Requester) {
		if r.Headers == nil {
			r.Headers = make(map[string]interface{})
		}
		r.Headers[key] = value
	}
}

func WithData(key, value string) RequestOption {
	return func(r *Requester) {
		if r.Data == nil {
			r.Data = make(map[string]interface{})
		}
		r.Data[key] = value
	}
}

func WithAuthorization(token string) RequestOption {
	return func(r *Requester) {
		r.Headers["Authorization"] = "Bearer " + token
	}
}

func WithResponse(resp interface{}) RequestOption {
	return func(r *Requester) {
		r.Response = resp
	}
}

func (r *Requester) Post(url string, options ...RequestOption) {
	for _, option := range options {
		option(r)
	}

	var requestBody []byte
	var err error
	var req *http.Request

	if requestBody, err = json.Marshal(&r.Data); err != nil {
		panic(err)
	}

	req, err = http.NewRequest(http.MethodPost, r.BaseUrl+url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	resp := r.make(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	r.decode(resp)
}

func (r *Requester) Get(url string, options ...RequestOption) {
	for _, option := range options {
		option(r)
	}

	var err error
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, r.BaseUrl+url, nil)
	if err != nil {
		panic(err)
	}

	resp := r.make(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	r.decode(resp)
}

func (r *Requester) make(req *http.Request) *http.Response {

	// set default timeout
	if r.Timeout == 0 {
		r.Timeout = time.Second * 5
	}

	// Add header
	for key, value := range r.Headers {
		req.Header.Add(key, value.(string))
	}

	// create http client
	client := &http.Client{
		Timeout: r.Timeout,
	}

	// make request
	if resp, err := client.Do(req); err != nil {
		panic(err)
	} else {
		return resp
	}
}

func (r *Requester) decode(resp *http.Response) {
	if resp != nil {
		decodeErr := json.NewDecoder(resp.Body).Decode(&r.Response)
		if decodeErr != nil {
			panic(decodeErr)
		}
	}
}
