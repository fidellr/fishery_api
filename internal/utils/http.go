package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	customError "github.com/pkg/errors"
)

type HttpClient struct {
	client http.Client
}

type Response struct {
	*http.Response
	Body []byte
}

type ApiError struct {
	Err string `json:"error"`
}

// Error interface implementation of golang Error
func (e ApiError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

func HandleHTTPError(statusCode int, body []byte) error {
	if statusCode >= 400 {
		var apiErr ApiError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return err
		}
		return apiErr
	}
	return nil
}

func (r *HttpClient) DoRequest(request *http.Request) (*Response, error) {
	resp, err := r.client.Do(request)
	if err != nil {
		return nil, customError.Wrap(err, "do request")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	// set response and close resp body
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, customError.Wrap(err, "read response body")
	}

	// return Response
	return &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}, nil
}

func (r *HttpClient) NewRequest(ctx context.Context, host, path string, req interface{}, method string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", host, path)
	var err error
	var body []byte
	var ok bool
	var request *http.Request

	switch method {
	case http.MethodPost:
		body, ok = req.([]byte)
		if !ok {
			body, err = json.Marshal(req)
			if err != nil {
				return nil, customError.Wrap(err, "marshal request body")
			}
		}
		request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, customError.Wrap(err, "new request post")
		}
	case http.MethodPut:
		body, ok = req.([]byte)
		if !ok {
			body, err = json.Marshal(req)
			if err != nil {
				return nil, customError.Wrap(err, "marshal request body")
			}
		}
		request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, customError.Wrap(err, "new request post")
		}
	case http.MethodGet:
		request, err = http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return nil, customError.Wrap(err, "new request get")
		}
	case http.MethodDelete:
		body, ok = req.([]byte)
		if !ok {
			body, err = json.Marshal(req)
			if err != nil {
				return nil, customError.Wrap(err, "marshal request body")
			}
		}
		request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
		fmt.Println(request.Body)
		if err != nil {
			return nil, customError.Wrap(err, "new request delete")
		}
	}

	request.Header.Set("Content-type", "application/json")
	return request, nil
}

func Request() *HttpClient {
	return &HttpClient{client: http.Client{Timeout: 5 * time.Second}}
}
