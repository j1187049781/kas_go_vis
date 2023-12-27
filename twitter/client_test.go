package twitter_test

import (
	"fmt"
	"kas_go_vis/twitter"
	"testing"
)

func TestTwitterClient_GetUserInfo(t *testing.T) {
	c := twitter.NewTwitterClient(twitter.BearerToken, twitter.Host, twitter.Proxy)
	res, err := c.GetUserInfo("USERNAME")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", res)
}

func TestGetFollowers(t *testing.T) {
	c := twitter.NewTwitterClient(twitter.BearerToken, twitter.Host, twitter.Proxy)
	res, err := c.GetFollowers("USERNAME")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", res)
}
