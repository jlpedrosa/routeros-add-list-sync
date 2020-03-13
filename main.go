package main

import (
	"flag"
	"log"

	"github.com/go-routeros/routeros"
	"github.com/jlpedrosa/routeros-add-list-sync/internal/list"
)

var (
	address  = flag.String("address", "10.40.10.2:8728", "RouterOS address and port")
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

	addListName:="private-addresses"
	err2, addList := lm.GetByName(addListName)
	if err2 != nil {
		log.Fatal(err2)
	}

	res := syncList([]string{}, addList)
	for _, line := range res.ToBeAdded {
		log.Print(line)
	}
}

type SyncActions struct {
	ToBeAdded []string
	ToBeDeleted []string
}

func syncList(actual,target []string) (SyncActions) {
	a := map[string]bool{}
	t := map[string]bool{}
	result := SyncActions {
	}

	for _, i := range actual { a[i]=true }
	for _, i := range target { t[i]=true }

	for _, i := range target { 
		if _ , ok := a[i]; !ok  {
			result.ToBeAdded = append(result.ToBeAdded, i)
		}		
	}

	for _, i := range actual { 
		if _ , ok := t[i]; !ok  {
			result.ToBeDeleted = append(result.ToBeAdded, i)
		}		
	}

	return result
}

