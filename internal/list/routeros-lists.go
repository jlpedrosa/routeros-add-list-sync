package list

import (
	"errors"
	"fmt"
	"net"

	"github.com/go-routeros/routeros"
)

type RouterOsListManager struct {
	client *routeros.Client
}

func NewRouterOsListManager(cli *routeros.Client) RouterOsListManager {
	return RouterOsListManager{
		client: cli,
	}
}

// GetV4ListByName returns the addresses in a ip v6 address list
func (lm *RouterOsListManager) GetV4ListByName(listName string) (error, []string) {
	command := "/ip/firewall/address-list/print"

	r, err := lm.client.RunArgs([]string{
		command,
		"?list=" + listName,
	})

	if err != nil {
		return err, nil
	}

	addressrange := []string{}
	for _, i := range r.Re {
		addressrange = append(addressrange, i.Map["address"])
	}

	return nil, addressrange
}

// GetV6ListByName returns the addresses in a ip v6 address list
func (lm *RouterOsListManager) GetV6ListByName(listName string) (error, []string) {
	command := "/ipv6/firewall/address-list/print"

	r, err := lm.client.RunArgs([]string{
		command,
		"?list=" + listName,
	})

	if err != nil {
		return err, nil
	}

	addressrange := []string{}
	for _, i := range r.Re {
		addressrange = append(addressrange, i.Map["address"])
	}

	return nil, addressrange
}

// AddAddressToList adds a route to a list.
func (lm *RouterOsListManager) AddAddressToList(address string, listName string) error {
	add, _, err := net.ParseCIDR(address)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to add address, not valid: %s", err))
	}

	if add.To4() != nil {
		return lm.AddIPV4AddressToList(address, listName)
	}

	return lm.AddIPV6AddressToList(address, listName)
}

// AddIPV4AddressToList adds a route to a list.
func (lm *RouterOsListManager) AddIPV4AddressToList(address string, listName string) error {
	command := "/ip/firewall/address-list/add"

	_, err := lm.client.RunArgs([]string{
		command,
		"=list=" + listName,
		"=address=" + address,
		"=comment=Generated",
	})

	if err != nil {
		return err
	}

	return nil
}

// AddIPV6AddressToList adds a route to a list.
func (lm *RouterOsListManager) AddIPV6AddressToList(address string, listName string) error {
	command := "/ipv6/firewall/address-list/add"

	_, err := lm.client.RunArgs([]string{
		command,
		"=list=" + listName,
		"=address=" + address,
		"=comment=Generated",
	})

	if err != nil {
		return err
	}

	return nil
}
