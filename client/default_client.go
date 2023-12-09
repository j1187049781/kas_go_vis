package client

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type DefaultClient struct {
	httpClient *http.Client
	host string
}

func NewDefaultClient() *DefaultClient {
	return &DefaultClient{
		httpClient: &http.Client{},
		host: "https://api-v2-do.kas.fyi",
	}
}

/**
Get https://api-v2-do.kas.fyi/analytics/addresses/richList 

[{"address":"kaspa:qq3k4du6wf2g26j7ds6fqmgtgavgm3zy676wntp2e52nsuns2n4s6xkndmx0y",
"balance":"92725114960945091",
"percent":4.275111932720819,
"tags":
[{"address":"kaspa:qq3k4du6wf2g26j7ds6fqmgtgavgm3zy676wntp2e52nsuns2n4s6xkndmx0y",
"name":"KuCoin Wallet","link":null}]}
]

**/

func (c *DefaultClient) GetRichTopList() ([]Address, error) {
	richAddrs, err:= c.getRichList()
	if err != nil {
		log.Printf("error getting rich list: %s", err)
		return nil, err
	}

	var addresses []Address
	for _, richAddr := range richAddrs {
		addr := richAddr.Address
		bStr := richAddr.Balance

		b, err := strconv.ParseFloat(bStr, 64)
		if err != nil {
			log.Printf("error parse float : %s", err)
			return nil, err
		}
		b_norm := b/100_000_000


		var addrTags []AddrTag
		for _, richTag := range richAddr.Tags{
			t := AddrTag(richTag)

			addrTags = append(addrTags, t)
		}

		addresses = append(addresses, Address{
			Address: addr,
			Balance: b_norm,
			Tags: addrTags,
		})
	}

	return addresses, nil
}

func (c *DefaultClient) GetAddrTag() ([]AddrTag, error) {
	richAddrs, err:= c.getRichList()
	if err != nil {
		log.Printf("error getting rich list: %s", err)
		return nil, err
	}

	var addrTags []AddrTag
	for _, richAddr := range richAddrs {
		if len(richAddr.Tags) == 0 {
			continue
		}

		for _, richTag := range richAddr.Tags{
			t := AddrTag(richTag)

			addrTags = append(addrTags, t)
		}
	}

	return addrTags, nil
}

func (c *DefaultClient) getRichList() ([]RichAddr, error){
	resp, err:= c.httpClient.Get(c.host + "/analytics/addresses/richList")
	if err != nil {
		log.Printf("error getting rich list: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	var richAddrs []RichAddr
	err = json.NewDecoder(resp.Body).Decode(&richAddrs)
	if err != nil {
		log.Printf("error decoding rich list: %s", err)
		return nil, err
	}

	return richAddrs, nil
}

type RichAddr struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Percent float64 `json:"percent"`
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Address string `json:"address"`
	Name string `json:"name"`
	Link string `json:"link"`
}