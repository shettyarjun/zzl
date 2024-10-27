package ip

import (
	"fmt"
	"net"
)

func GenerateIPs(startIP, endIP string) ([]string, error) {
	var ips []string

	// Parse the start and end IPs
	start := net.ParseIP(startIP).To4()
	end := net.ParseIP(endIP).To4()
	if start == nil || end == nil {
		return nil, fmt.Errorf("invalid IP address")
	}

	// Loop from start to end
	for ip := start; !ip.Equal(end); incrementIP(ip) {
		ips = append(ips, ip.String())
	}
	// Add the end IP to the list
	ips = append(ips, end.String())

	return ips, nil
}

func GenerateEndIP(startIP string) (string, error) {
	ip := net.ParseIP(startIP).To4()
	if ip == nil {
		return "", fmt.Errorf("invalid IPv4 address")
	}

	ip[3] = 255 // Set the last octet to 255
	return ip.String(), nil
}

// incrementIP modifies the IP address to the next address
// this doesn't need to be exposed, keep it like that
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}