package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/logging"

	"github.com/geovanisouza92/terraform-provider-dreamfactory/dreamfactory/types"
)

// admin app API key
const apiKey = "6498a8ad1beb9d84d63035c5d1120c007fad6de706734db9689f8996707e0f7d"

// Client is used to consume DreamFactory's system/ API
type Client struct {
	endpoint string
	hc       *http.Client
	session  *types.Session
}

// New creates a new Client instance, of an error if login credentials are invalid
func New(endpoint, email, password string, hc *http.Client) (*Client, error) {
	c := Client{
		endpoint: endpoint,
		hc:       hc,
	}
	s, err := c.login(email, password)
	if err != nil {
		return nil, err
	}
	c.session = &s
	return &c, err
}

func (c *Client) login(email, password string) (s types.Session, err error) {
	l := map[string]string{"email": email, "password": password}
	err = c.send("POST", "/api/v2/system/admin/session", 200, l, &s)
	return
}

func (c *Client) send(method, path string, expectedStatusCode int, in, out interface{}) error {
	var br io.ReadWriter

	// Serialize request body
	if in != nil {
		br = &bytes.Buffer{}
		if err := json.NewEncoder(br).Encode(in); err != nil {
			return err
		}

		if logging.IsDebugOrHigher() {
			buf := br.(*bytes.Buffer)
			log.Println("Request: " + method + " " + path + "\n" + string(buf.Bytes()))
		}
	}

	// Create request
	req, err := http.NewRequest(method, c.endpoint+path, br)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-DreamFactory-Api-Key", apiKey)
	if c.session != nil {
		req.Header.Add("X-DreamFactory-Session-Token", c.session.SessionToken)
	}

	// Send request
	res, err := c.hc.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return err
	}

	// Handle unexpected error
	if res.StatusCode != expectedStatusCode {
		b, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("api response with status: %d\n%s", res.StatusCode, string(b))
	}

	// Ignore response?
	if out == nil {
		return nil
	}

	b, _ := ioutil.ReadAll(res.Body)
	buf := bytes.NewBuffer(b)

	if logging.IsDebugOrHigher() {
		log.Println("Response: " + method + " " + path + "\n" + string(buf.Bytes()))
	}

	// Parse response
	if err = json.NewDecoder(buf).Decode(out); err != nil {
		return err
	}

	return nil
}
