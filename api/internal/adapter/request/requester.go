package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Requester struct {
	client http.Client
}

func NewRequester() *Requester {
	return &Requester{client: http.Client{}}
}

func (r *Requester) into(body io.ReadCloser, into any) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(into)
}

func (r *Requester) request(method, url, access string) (io.ReadCloser, error) {
	req, _ := http.NewRequest(method, url, nil)

	if access != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access))
	}

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("status code not 200")
	}
	return response.Body, nil
}

func (r *Requester) Get(url string) (io.ReadCloser, error) {
	return r.request("GET", url, "")
}

func (r *Requester) PostJSON(url string, values url.Values, into any) error {
	res, err := http.PostForm(url, values)
	if err != nil {
		return err
	}
	return r.into(res.Body, into)
}

func (r *Requester) GetJSONAuthenticated(url string, access string, into any) error {
	body, err := r.request("GET", url, "")
	if err != nil {
		return err
	}
	return r.into(body, into)
}

func (r *Requester) GetJSON(url string, into any) error {
	return r.GetJSONAuthenticated(url, "", into)
}
