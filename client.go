package top

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// TopClient 客户端
type TopClient struct {
	appKey     string
	appSecret  string
	url        string
	signMethod SignMethod

	httpClient *http.Client
	logger     LoggerInterface
}

var (
	productionTestURL = "https://eco.taobao.com/router/rest"
	productionURL     = "https://eco.taobao.com/router/rest"
	overseasURL       = "https://api.taobao.com/router/rest"

	APIVersion = "2.0"
)

const (
	EnvOverSeas       = "overseas"
	EnvProduction     = "production"
	EnvProductionTest = "production_test"
)

func NewTopClient(appKey string, appSecret string, options ...TopClientOption) *TopClient {
	tc := &TopClient{
		httpClient: &http.Client{},
		appKey:     appKey,
		appSecret:  appSecret,
		url:        productionURL,
		signMethod: SignMethodHMAC,
	}
	for _, option := range options {
		option(tc)
	}
	return tc
}

type TopClientOption func(*TopClient)

func Logger(logger LoggerInterface) TopClientOption {
	return func(c *TopClient) {
		c.logger = logger
	}
}
func HttpClient(client *http.Client) TopClientOption {
	return func(c *TopClient) {
		c.httpClient = client
	}
}
func SignatureMethod(signMethod SignMethod) TopClientOption {
	return func(c *TopClient) {
		c.signMethod = signMethod
	}
}

func Env(env string) TopClientOption {
	return func(c *TopClient) {
		if env == EnvProduction {
			c.url = productionURL
		} else if env == EnvProductionTest {
			c.url = productionTestURL
		} else {
			c.url = overseasURL
		}
	}
}

func (c *TopClient) execute(ctx context.Context, req TopRequest) ([]byte, error) {
	if isNil(req) {
		return nil, ErrorRequestIsNil
	}
	err := req.Signature(ctx, c.appKey, c.appSecret, c.signMethod)
	if err != nil {
		return nil, fmt.Errorf("req.Signature error %w", err)
	}

	reqData := getRequestData(req)

	log.Printf("request: %s\n", reqData)
	reqBytes := bytes.NewReader(reqData)

	httpReq, err := http.NewRequest(http.MethodPost, c.url, reqBytes)
	if err != nil {
		return nil, fmt.Errorf("new request failed %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request faield %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response faield %w", err)
	}
	log.Printf("response %s\n", respBytes)
	return respBytes, nil
}
