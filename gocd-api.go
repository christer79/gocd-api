package gocdapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "http://localhost:8153/go/"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	Agents        *AgentService
	Users         *UserService
	PipelineGroup *PipelineGroupService
	Artifacts     *ArtifactService
}

// Response is a GitHub API response.  This wraps the standard http.Response
// returned from GitHub.
type Response struct {
	*http.Response
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	fmt.Println("baseurl: " + baseURL.String())

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.Agents = &AgentService{client: c}
	c.Artifacts = &ArtifactService{client: c}
	c.Users = &UserService{client: c}
	c.PipelineGroup = &PipelineGroupService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("christer.eriksson", "p4rqBE24")

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	fmt.Println("URL:" + req.RequestURI)
	fmt.Println("URL:" + req.RemoteAddr)
	resp, err := c.client.Do(req)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	fmt.Println("Status: " + resp.Status)

	defer resp.Body.Close()
	response := &Response{Response: resp}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			/*
									buf := new(bytes.Buffer)
									buf.ReadFrom(resp.Body)
									s := buf.String()
				  				fmt.Println("BODY: " + s)
			*/
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				fmt.Println("Error on dECODE" + err.Error())
				fmt.Println(err)
			}
		}
	}
	return response, err
}
