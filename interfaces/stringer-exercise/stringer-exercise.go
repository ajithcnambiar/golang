package main

import "fmt"

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[3], ipaddr[2], ipaddr[1], ipaddr[0])
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
