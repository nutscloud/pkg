package net

var Bridge Bridger

type Bridger interface {
	GetBridgeInterfaces(br string) ([]string, error)
	GetAllBridge() ([]string, error)
}
