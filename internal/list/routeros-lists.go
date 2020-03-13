package list

import (
	"github.com/go-routeros/routeros"
)

type ListManager interface {
	GetByName()
}


type RouterOsListManager struct {
	client *routeros.Client 
}

func NewRouterOsListManager(cli *routeros.Client) RouterOsListManager {
	return RouterOsListManager {
		client: cli,
	}
}
 
func (lm *RouterOsListManager) GetByName(listName string) (error, []string) {
	command :=  "/ip/firewall/address-list/print" //?list=NordVPN-USA-SRC-LIST
	
	r, err := lm.client.RunArgs([]string { 
			command,
			"?list="+listName ,
	})
	
	if err != nil {
		return err, nil
	}

	addressrange := []string{}
	for _, i := range r.Re {
		addressrange = append(addressrange, i.Map["address"])
	}

	return  nil, addressrange
}