package main

import (
	"fmt"
)

// Exercise: Stringers
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	// First method.
	// var strArray []string
	// for _, byteItem := range ipAddr {
	// 	strArray = append(strArray, fmt.Sprintf("%d", byteItem))
	// }
	// res := strings.Join(strArray, ".")
	// - - -

	// Second method.
	res := ""
	for _, byteItem := range ipAddr {
		res += fmt.Sprintf("%d.", byteItem)
	}
	res = res[:len(res)-1]
	// - - -
	return res
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
