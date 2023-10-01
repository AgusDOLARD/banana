package main

import (
	"flag"
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

var areIPs bool
var useClipboard bool

func init() {
	flag.BoolVar(&areIPs, "ip", false, "son ips (:")
	flag.BoolVar(&useClipboard, "c", false, "usa el portapapeles")
	flag.Parse()
}

func main() {
	var entries []string

	if useClipboard {
		cbtext := readFromClipboard()
		entries = strings.Split(cbtext, "\n")
	} else {
		entries = flag.Args()
	}

	for _, e := range entries {
		if e != "" {
			if areIPs {
				spitSubnet(e)
			} else {
				spitFqdn(e)
			}

		}
	}
}

func readFromClipboard() string {
	cb := clipboard.Read(clipboard.FmtText)
	return string(cb)
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
