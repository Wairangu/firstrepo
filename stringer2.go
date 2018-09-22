package main

import (
	"fmt"
)

//IPAddr allows fmt.Printf to print ip addresses as decimals
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", int(ip[0]), int(ip[1]), int(ip[3]), int(ip[3]))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
