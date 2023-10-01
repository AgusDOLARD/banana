package main

import (
	"flag"
	"fmt"
)

var areIPs bool

func init() {
	flag.BoolVar(&areIPs, "ip", false, "son ips (:")
	flag.Parse()
}

func main() {
	entries := flag.Args()

	for _, e := range entries {
		if areIPs {
			spitSubnet(e)
		} else {
			spitFqdn(e)
		}
	}
}

func spitFqdn(domain string) {
	fmt.Printf("set type fqdn\n")
	fmt.Printf("set fqdn \"%s\"\n", domain)
	fmt.Printf("next\n\n")
}

func spitSubnet(ip string) {
	fmt.Printf("edit \"BL_%s\"\n", ip)
	fmt.Printf("set subnet \"%s/32\"\n", ip)
	fmt.Printf("next\n\n")
}
