package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var areIPs bool

func init() {
	flag.BoolVar(&areIPs, "ip", false, "son ips (:")
	flag.Parse()
}

func main() {
	var entries []string
	if len(flag.Args()) != 0 {
		entries = flag.Args()
	} else {
		entries = readFromCli()
	}

	for _, e := range entries {
		if areIPs {
			spitSubnet(e)
		} else {
			spitFqdn(e)
		}
	}
}

func readFromCli() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
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
