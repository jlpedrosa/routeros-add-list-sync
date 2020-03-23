package asnfetcher

import "testing"

func TestFetcherAsOverview(t *testing.T) {
	res, err := AsOverview("2906")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestFetcherAsAdvertisement(t *testing.T) {
	res, err := AnnouncedPrefixes("2906")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
