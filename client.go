package n2yo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Client struct {
	c       *http.Client
	apiKey  string
	baseURL string
}

func New(apiKey string) (c N2YOer) {
	client := &Client{}
	client.apiKey = apiKey
	client.c = &http.Client{}
	client.c.Timeout = DefaultClientTimeout
	client.baseURL = BaseURL

	c = client
	return
}

func (c *Client) CustomHTTPClient(cl *http.Client) {
	c.c = cl
}

func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

func (c *Client) GetTLE(id int) (r Response, err error) {
	var reqURL *url.URL
	reqURL, err = url.Parse(c.baseURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse base url")
		return
	}
	reqURL.Path = fmt.Sprintf(TLEPathFormat, id)
	q := reqURL.Query()
	q.Add(APIKeyQuery, c.apiKey)
	reqURL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = c.c.Get(reqURL.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to send get request to %s", reqURL.String())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("get tle request to %s failed with status %d", reqURL.String(), resp.StatusCode)
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read get tle response body")
		return
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal get tle response")
		return
	}

	if r.Error != "" {
		err = fmt.Errorf("error response %s", r.Error)
		return
	}

	return
}

func (c *Client) GetPositions(id int, obsLat, obsLang, obsAlt float64, seconds int) (r Response, err error) {
	var reqURL *url.URL
	reqURL, err = url.Parse(c.baseURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse base url")
		return
	}
	reqURL.Path = fmt.Sprintf(PositionsPathFormat, id, obsLat, obsLang, obsAlt, seconds)
	q := reqURL.Query()
	q.Add(APIKeyQuery, c.apiKey)
	reqURL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = c.c.Get(reqURL.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to send get request to %s", reqURL.String())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("get positions request to %s failed with status %d", reqURL.String(), resp.StatusCode)
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read get positions response body")
		return
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal get positions response")
		return
	}

	if r.Error != "" {
		err = fmt.Errorf("error response %s", r.Error)
		return
	}

	return
}

func (c *Client) GetVisualPasses(id int, obsLat, obsLang, obsAlt float64, days, minVisibility int) (r Response, err error) {
	var reqURL *url.URL
	reqURL, err = url.Parse(c.baseURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse base url")
		return
	}
	reqURL.Path = fmt.Sprintf(VisualPassesPathFormat, id, obsLat, obsLang, obsAlt, days, minVisibility)
	q := reqURL.Query()
	q.Add(APIKeyQuery, c.apiKey)
	reqURL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = c.c.Get(reqURL.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to send get request to %s", reqURL.String())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("get visual passes request to %s failed with status %d", reqURL.String(), resp.StatusCode)
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read get visual passes response body")
		return
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal get visual passes response")
		return
	}

	if r.Error != "" {
		err = fmt.Errorf("error response %s", r.Error)
		return
	}

	r.PassesType = PassTypeVisual

	return
}

func (c *Client) GetRadioPasses(id int, obsLat, obsLang, obsAlt float64, days, minElevation int) (r Response, err error) {
	var reqURL *url.URL
	reqURL, err = url.Parse(c.baseURL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse base url")
		return
	}
	reqURL.Path = fmt.Sprintf(RadioPassesPathFormat, id, obsLat, obsLang, obsAlt, days, minElevation)
	q := reqURL.Query()
	q.Add(APIKeyQuery, c.apiKey)
	reqURL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = c.c.Get(reqURL.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to send get request to %s", reqURL.String())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		err = fmt.Errorf("get radio passes request to %s failed with status %d", reqURL.String(), resp.StatusCode)
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read get radio passes response body")
		return
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal get radio passes response")
		return
	}

	if r.Error != "" {
		err = fmt.Errorf("error response %s", r.Error)
		return
	}

	r.PassesType = PassTypeRadio

	return
}
