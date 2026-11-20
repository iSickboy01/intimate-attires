package main

import (
	"fmt"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHTTPClient(client *http.Client) *HttpClient {
	return &HttpClient{
		client: client,
	}
}

func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "intimateattires-API/1.0")
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "image/jpeg")
	}

	fmt.Printf("Sending %s request to %s\n", req.Method, req.URL)

	resp, err := h.client.Do(req)

	if err != nil {

		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	fmt.Printf("Received status: %d\n", resp.StatusCode)
	return resp, nil

}
