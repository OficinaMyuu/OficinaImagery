package service

import ipdata "github.com/ipdata/go"

var ipDataClient *ipdata.Client

func SetIpDataClient(c *ipdata.Client) {
	ipDataClient = c
}

func IsSecure(ip string) (bool, *APIError) {
	result, err := ipDataClient.Lookup(ip)
	if err != nil {
		return nil,
	}
}
