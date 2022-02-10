package net

import (
	"net"
	"reflect"
	"sort"
	"testing"
)

type nic struct {
	name    string
	ipaddrs []string
}

// TestGetInterfaceAddrs test GetInterfaceAddrs
func TestGetInterfaceAddrs(t *testing.T) {
	nics, err := getInterfacesAddrs()
	if err != nil {
		t.Fatalf("[getInterfacesAddrs error]")
	}

	for _, n := range nics {
		addrs, err := GetInterfaceAddrs(n.name)
		if err != nil {
			t.Fatalf("[GetInterfaceAddrs error]: %v", err)
		}
		// The effect of removal order on results.
		sort.Strings(addrs)
		sort.Strings(n.ipaddrs)

		if !reflect.DeepEqual(addrs, n.ipaddrs) {
			t.Fatalf("actual addrs was '%v'; "+
				"want '%v'", addrs, n.ipaddrs)
		}

	}
}

// get all nics name and ipaddress
func getInterfacesAddrs() ([]*nic, error) {
	nics := []*nic{}

	ifs, err := net.Interfaces()
	if err != nil {
		return nics, err
	}

	for _, i := range ifs {
		_nic := &nic{name: i.Name}
		addrs, err := i.Addrs()
		if err != nil {
			return nics, err
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() &&
				ipnet.IP.To4() != nil {
				_nic.ipaddrs = append(_nic.ipaddrs, ipnet.IP.String())
			}
		}
		nics = append(nics, _nic)
	}

	return nics, nil
}
