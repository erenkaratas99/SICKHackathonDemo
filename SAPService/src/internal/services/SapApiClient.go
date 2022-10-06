package services

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io/ioutil"
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

func (rc *RestClient) DoGetReqToSAPCloud(URI string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	req.Header.Set("APIKey", "Xos0AuVGGL8jAVmqjM2MGLAIZi3lxw4k") // I'm so sick of that solution too
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()
	fasthttp.ReleaseResponse(res)
	if err := rc.Client.Do(req, res); err != nil {
		return err
	}
	var body []byte
	if bytes.EqualFold(res.Header.Peek("Content-Encoding"), []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = res.BodyGunzip()
	} else {
		body = res.Body()
	}

	fmt.Println("body : ", string(body), "\n\n")
	err := ioutil.WriteFile("output.xml", body, 0644)

	if err != nil {
		log.Info(err)
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return nil
}
