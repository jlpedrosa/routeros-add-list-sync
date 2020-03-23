package main

import (
	"flag"
	"log"
	"net"

	"github.com/go-routeros/routeros"
	asnfetcher "github.com/jlpedrosa/routeros-add-list-sync/internal/asnfecher"
	"github.com/jlpedrosa/routeros-add-list-sync/internal/list"
)

var (
	address  = flag.String("address", "10.40.20.2:8728", "RouterOS address and port")
	username = flag.String("username", "jose", "User name")
)

func dial() (*routeros.Client, error) {

	return routeros.Dial(*address, *username, *password)
}

func main() {
	flag.Parse()

	c, err := dial()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	lm := list.NewRouterOsListManager(c)

	asn := []string{"2906", "40027", "55095"}

	fetcher := asnfetcher.NewFetcher(asn)

	allAdresses, _ := fetcher.Fetch()
	allAdresses = discardDuplicated(allAdresses)
	ipv4asnadds := []string{}
	ipv6asnadds := []string{}

	for _, addStr := range allAdresses {
		add, _, err := net.ParseCIDR(addStr)
		if err != nil {
			log.Fatal("Invalid IP")
		}

		if add.To4() != nil {
			ipv4asnadds = append(ipv4asnadds, addStr)
		} else {
			ipv6asnadds = append(ipv6asnadds, addStr)
		}
	}

	addListName := "netlix"
	err2, addListV4 := lm.GetV4ListByName(addListName)
	if err2 != nil {
		log.Fatal(err2)
	}

	err2, addListV6 := lm.GetV6ListByName(addListName)
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Print("Syncing router")
	res := syncList(addListV4, ipv4asnadds)
	for _, line := range res.ToBeAdded {
		log.Printf("Adding prefix:%s", line)
		if err := lm.AddAddressToList(line, addListName); err != nil {
			log.Fatalf("Error adding to list %s", err)
		}
	}

	res = syncList(addListV6, ipv6asnadds)
	for _, line := range res.ToBeAdded {
		log.Printf("Adding prefix:%s", line)
		if err := lm.AddAddressToList(line, addListName); err != nil {
			log.Fatalf("Error adding to list %s", err)
		}
	}

}

// SyncActions Describes the changes to take place in two sets to sync them.
type SyncActions struct {
	ToBeAdded   []string
	ToBeDeleted []string
}

func discardDuplicated(list []string) []string {
	set := map[string]bool{}

	for _, item := range list {
		set[item] = true
	}

	keys := []string{}
	for key, _ := range set {
		keys = append(keys, key)
	}

	return keys
}

func syncList(actual, target []string) SyncActions {
	a := map[string]bool{}
	t := map[string]bool{}
	result := SyncActions{}

	for _, i := range actual {
		a[i] = true
	}
	for _, i := range target {
		t[i] = true
	}

	for _, i := range target {
		if _, ok := a[i]; !ok {
			result.ToBeAdded = append(result.ToBeAdded, i)
		}
	}

	for _, i := range actual {
		if _, ok := t[i]; !ok {
			result.ToBeDeleted = append(result.ToBeAdded, i)
		}
	}

	return result
}
