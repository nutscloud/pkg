package net

import (
	"errors"
	"net"

	"github.com/nutscloud/go-util/math"
)

// IPString2Int Converts the IP string to a number
func IPString2Uint(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// NetIP2Uint Converts the net.IP to a number
func NetIP2Uint(ip net.IP) uint {
	ipv4 := ip.To4()
	return uint(ipv4[3]) | uint(ipv4[2])<<8 | uint(ipv4[1])<<16 | uint(ipv4[0])<<24
}

func IPRangeIsOverlap(aStart, aEnd, bStart, bEnd net.IP) bool {
	aStartInt := NetIP2Uint(aStart)
	aEndInt := NetIP2Uint(aEnd)
	bStartInt := NetIP2Uint(bStart)
	bEndInt := NetIP2Uint(bEnd)

	return math.IsRangeOverlapUint(aStartInt, aEndInt, bStartInt, bEndInt)
}

type IPs []string

func (ips IPs) Len() int { return len(ips) }
func (ips IPs) Less(i, j int) bool {
	ipi := net.ParseIP(ips[i]).To4()
	ipj := net.ParseIP(ips[j]).To4()

	for _i := range ipi {
		switch {
		case ipi[_i] == ipj[_i]:
			continue
		case ipi[_i] < ipj[_i]:
			return true
		default:
			return false
		}
	}
	return false
}
func (ips IPs) Swap(i, j int) { ips[i], ips[j] = ips[j], ips[i] }
