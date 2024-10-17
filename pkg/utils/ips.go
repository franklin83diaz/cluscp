package utils

import (
	"math/big"
	"net"
)

// Receive a string with the range of ips
// example: "10.0.0.1-3"
// and return the List of ips
// example: ["10.0.0.1", "10.0.0.2", "10.0.0.3"]
func GetIpList(s string) {
	//!TODO: Implement GetIpList
}

// Convert an IP to an int
func ipToInt(ip net.IP) *big.Int {
	ip = ip.To4()
	return big.NewInt(0).SetBytes(ip)
}

// Convert an int to an IP
func intToIP(ipInt *big.Int) net.IP {
	fourBytes := make([]byte, 4)
	ipInt.FillBytes(fourBytes)
	return net.IP(fourBytes).To4()
}
