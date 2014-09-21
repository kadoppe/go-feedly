package feedly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://cloud.feedly.com/"
	apiVersion     = "v3"
	userAgent      = "go-feedly" + libraryVersion
)

type Client struct {
	client     *http.Client
	BaseURL    *url.URL
	UserAgent  string
	Categories *CategoriesService
	Profile    *ProfileService
	Markers    *MarkersService
	// Entries       *EntriesService
	// Feeds         *FeedsService
	// Mixes         *MixesService
	// Search        *SearchService
	// Streams       *StreamsService
	// Subscriptions *SubscriptionsService
	// Tags          *TagsService
	// Topics        *TopicsService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.Categories = &CategoriesService{client: c}
	c.Profile = &ProfileService{client: c}
	c.Markers = &MarkersService{client: c}
	// c.Entries = &EntriesService{client: c}
	// c.Feeds = &FeedsService{client: c}
	// c.Mixes = &MixesService{client: c}
	// c.Search = &SearchService{client: c}
	// c.Streams = &StreamsService{client: c}
	// c.Subscriptions = &SubscriptionsService{client: c}
	// c.Tags = &TagsService{client: c}
	// c.Topics = &TopicsService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(apiVersion + "/" + urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
