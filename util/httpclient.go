package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DavidHsaiou/dcom/errors"
)

type HttpClient interface {
	Get(url string) (string, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() HttpClient {
	return &httpClient{
		client: &http.Client{},
	}
}

func (h *httpClient) Get(url string) (string, error) {
	resp, err := h.internalCall(http.MethodGet, url, nil, "")
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (h *httpClient) internalCall(method string, url string, header map[string]string, body string) (string, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			DefaultLogger.Error(err)
		}
	}(resp.Body)

	if !h.checkStatusSuccess(resp.StatusCode) {
		return "", errors.HttpNotSuccess
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return string(responseBody), nil
}

func (h *httpClient) checkStatusSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
