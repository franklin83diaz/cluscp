package utils

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

// Receive a string with the range of ips
// example: "10.0.0.1-3"
// and return the List of ips
// example: ["10.0.0.1", "10.0.0.2", "10.0.0.3"]
func GetIpList(s string) ([]string, error) {
	ipPartStartSplit := make([]string, 4)
	ipPartEndSplit := make([]string, 4)

	// Split the string in two parts ipStart and ipEnd
	ips := strings.TrimSpace(s)
	ipsStartEndSplit := strings.Split(ips, "-")
	// Check if the ip range is valid
	if len(ipsStartEndSplit) != 2 {
		return nil, fmt.Errorf("invalid ip range")
	}

	//set the ipStart
	temp := strings.Split(ipsStartEndSplit[0], ".")
	// Check if the ipStart is valid
	if len(temp) != 4 {
		return nil, fmt.Errorf("invalid ip range")
	}
	copy(ipPartStartSplit, temp)

	//set the ipEnd
	ipPartEnd := strings.Split(ipsStartEndSplit[1], ".")

	// Set the End Ip String
	lenPartEnd := len(ipPartEnd)

	// Check if the ipEnd is valid
	if lenPartEnd > 4 {
		return nil, fmt.Errorf("invalid ip range")
	}

	switch lenPartEnd {
	case 1:
		ipPartEndSplit[0] = ipPartStartSplit[0]
		ipPartEndSplit[1] = ipPartStartSplit[1]
		ipPartEndSplit[2] = ipPartStartSplit[2]
		ipPartEndSplit[3] = ipPartEnd[0]
	case 2:
		ipPartEndSplit[0] = ipPartStartSplit[0]
		ipPartEndSplit[1] = ipPartStartSplit[1]
		ipPartEndSplit[2] = ipPartEnd[0]
		ipPartEndSplit[3] = ipPartEnd[1]
	case 3:
		ipPartEndSplit[0] = ipPartStartSplit[0]
		ipPartEndSplit[1] = ipPartEnd[0]
		ipPartEndSplit[2] = ipPartEnd[1]
		ipPartEndSplit[3] = ipPartEnd[2]
	case 4:
		ipPartEndSplit = ipPartEnd
	default:
		return nil, fmt.Errorf("invalid ip range")
	}

	ipStartStr := strings.Join(ipPartStartSplit, ".")
	ipStart := net.ParseIP(ipStartStr)
	if ipStart == nil {
		return nil, fmt.Errorf("invalid ip range")
	}
	ipStartInt := ipToInt(ipStart)

	ipEndStr := strings.Join(ipPartEndSplit, ".")
	ipEnd := net.ParseIP(ipEndStr)
	if ipEnd == nil {
		return nil, fmt.Errorf("invalid ip range")
	}
	ipEndInt := ipToInt(ipEnd)

	var ipList []string

	for i := ipStartInt; i.Cmp(ipEndInt) <= 0; i.Add(i, big.NewInt(1)) {
		ipList = append(ipList, intToIP(i).String())
	}

	return ipList, nil
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
