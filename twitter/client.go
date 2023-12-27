package twitter

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

var ApiKey = "bcHua8INAnQjNVHZVhz7MERBV"
var ApiSerect = "SaCS8DvHA2v88oaEl0ysm6YHH3kZlTkgVrT0mOh2M5Icbs8YyK"
var BearerToken = "AAAAAAAAAAAAAAAAAAAAAO2VrgEAAAAAP7ZFbfU54Z5G9y2CMPxWtbG3mAQ%3DJcmHb33hs3pB17jqkhrRexlyM7yiTrnAxskJtUFhS0eWRQH72n"

var AccessToken = "802132034109861888-nP9wDKU72sRgHMjbZBJy17fvkz8Kz1u"
var AccessTokenSerect = "N2CjeTsGLHhssD6j0LFvIkF6gPxeR0sfpSpAHGKhWhlKX"

var Host = "https://api.twitter.com"
var Proxy = "http://192.168.110.50:7890"

type TwitterClient struct {
	Token  string
	Host   string
	client *http.Client
}

func NewTwitterClient(token, host, proxy string) *TwitterClient {
	return &TwitterClient{
		Token: token,
		Host:  host,
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: func(r *http.Request) (*url.URL, error) {
					return url.Parse(proxy)
				},
			},
		},
	}
}

/*
*

	GET https://api.twitter.com/2/users/by/username/USER_NAME
*/
func (t *TwitterClient) GetUserInfo(username string) (map[string]interface{}, error) {
	url := t.Host + "/2/users/by/username/" + username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+t.Token)
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
GET https://api.twitter.com/2/users/:id/followers
*/
func (t *TwitterClient) GetFollowers(username string) (map[string]interface{}, error) {
	url := t.Host + "/2/users/" + username + "/followers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+t.Token)
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
