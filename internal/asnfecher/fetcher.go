package asnfetcher

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Fetcher Gets information from stat.ripe.net
type Fetcher struct {
	ASNs []string
}

// NewFetcher Creates a new AS information fetcher
func NewFetcher(ASlist []string) *Fetcher {
	return &Fetcher{
		ASNs: ASlist,
	}

}

// Fetch retrieves all the address ranges for the ASNs
func (f *Fetcher) Fetch() ([]string, error) {

	addr := []string {}

	for _, i := range f.ASNs {
		res, err := AnnouncedPrefixes(i)
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range res.Data.Prefixes {
			addr = append(addr, j.Prefix)
		}
		
	}

	return addr, nil 

}

// AsOverview Gets general information about an AS
func AsOverview(asn string) (*RipeAsOverviewResponse, error) {
	fetchURL := "https://stat.ripe.net/data/as-overview/data.json?resource=%s"
	url := fmt.Sprintf(fetchURL, asn)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var result RipeAsOverviewResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AnnouncedPrefixes Gets general information about an AS
func AnnouncedPrefixes(asn string) (*RipeAnnouncedPrefixesResponse, error) {
	fetchURL := "https://stat.ripe.net/data/announced-prefixes/data.json?soft_limit=ignore&resource=AS%s"
	url := fmt.Sprintf(fetchURL, asn)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var result RipeAnnouncedPrefixesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
