package net

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	SYSCLASSNET = "/sys/class/net"
)

type linuxBridge struct{}

func (_ *linuxBridge) GetBridgeInterfaces(br string) ([]string, error) {
	ifaces := make([]string)
	brIfPath := filepath.Join(SYSCLASSNET, br, "brif")

	ifs, err := ioutil.ReadDir(brIfPath)
	if err != nil {
		return ifaces, err
	}
	for _, i := range ifs {
		ifaces = append(ifaces, i.Name())
	}

	return ifaces, nil
}

func (_ *linuxBridge) GetAllBridge() ([]string, error) {
	bridge := make([]string)

	devs, err := ioutil.ReadDir(SYSCLASSNET)
	if err != nil {
		return bridge, err
	}

	for _, d := range devs {
		bridgePath := filepath.Join(SYSCLASSNET, d.Name(), "bridge")
		stat, err := os.Stat(bridgePath)
		if err != nil && os.IsNotExist(err) {
			continue
		} else {
			return bridge, err
		}

		if stat.IsDir() {
			bridge = append(bridge, d.Name())
		}
	}

	return bridge, nil
}
