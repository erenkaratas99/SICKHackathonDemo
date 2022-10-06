package services

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

type RestClient struct {
	Client *fasthttp.Client
}

var SingleRestClient *RestClient

func NewSingletonClient() *RestClient {
	SingleRestClient = &RestClient{Client: &fasthttp.Client{}}
	return SingleRestClient
}

func (c *RestClient) MakePostRequest(URI string, reqBody []byte) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(URI)
	req.Header.SetMethod(http.MethodPost)
	req.SetBody(reqBody)
	req.Header.SetContentType("application/json")
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := c.Client.Do(req, res); err != nil {
		return err
	}
	return nil
}
