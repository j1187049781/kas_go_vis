package client_test

import (
	"kas_go_vis/client"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	c := client.NewDefaultClient()

	richList, err := c.GetRichTopList()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for _, v := range richList {
		t.Logf("rich addr: %v", v)
	}

	richTag, err := c.GetAddrTag()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for _, v := range richTag {
		t.Logf("rich tag: %v", v)
	}
}