package net

import (
	"net"
)

// GetInterfaceAddrs get NIC ip address by NIC name
func GetInterfaceAddrs(ifname string) (addrs []string, err error) {
	i, err := net.InterfaceByName(ifname)
	if err != nil {
		return
	}

	addrlist, err := i.Addrs()
	if err != nil {
		return
	}

	for _, a := range addrlist {
		if ipnet, ok := a.(*net.IPNet); ok &&
			!ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			addrs = append(addrs, ipnet.IP.String())
		}
	}
	return
}
