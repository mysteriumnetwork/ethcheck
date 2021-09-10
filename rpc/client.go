package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Method = "POST"
const HttpContentType = "application/json"

const discardLimit int64 = 128 * 1024

type HTTPRPCClient struct {
	endpoint   *url.URL
	httpClient *http.Client
}

type Client interface {
	Call(ctx context.Context, resDst interface{}, method string, args ...interface{}) error
}

func NewHTTPRPCClient(endpoint *url.URL, httpClient *http.Client) *HTTPRPCClient {
	if endpoint == nil {
		panic(errors.New("nil endpoint received by RPC client constructor"))
	}

	if httpClient == nil {
		httpClient = NewDefaultHTTPClient()
	}

	return &HTTPRPCClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

func NewDefaultHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return errors.New("no redirects allowed")
		},
	}
}

func (c *HTTPRPCClient) Call(ctx context.Context, resDst interface{}, method string, args ...interface{}) error {
	req, err := NewUniqueRequest(method, args)
	if err != nil {
		return err
	}

	var reqBuffer bytes.Buffer
	reqEncoder := json.NewEncoder(&reqBuffer)
	err = reqEncoder.Encode(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, Method, c.endpoint.String(), &reqBuffer)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", HttpContentType)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return fmt.Errorf("RPC Call (%s, %v): bad status code %d", method, args, httpResp.StatusCode)
	}

	defer cleanupBody(httpResp.Body)
	decoder := json.NewDecoder(httpResp.Body)
	var resp Response
	err = decoder.Decode(&resp)
	if err != nil {
		return err
	}

	if resp.Version != req.Version {
		return errors.New("RPC version mismatch")
	}
	if resp.ID != req.ID {
		return errors.New("request ID mismatch")
	}

	return resp.UnmarshalResult(resDst)
}

// Does cleanup of HTTP response in order to make it reusable by keep-alive
// logic of HTTP client
func cleanupBody(body io.ReadCloser) {
	io.Copy(ioutil.Discard, &io.LimitedReader{
		R: body,
		N: discardLimit,
	})
	body.Close()
}
