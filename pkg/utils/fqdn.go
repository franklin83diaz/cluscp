package utils

import "fmt"

// GetFqdnList returns a list of FQDNs
// Receives the hostname, start and end of the counter and the domain
func GetFqdnList(hostname string, start int, end int, domain string) ([]string, error) {
	listFqdn := make([]string, 0)

	if start > end {
		return nil, fmt.Errorf("start is greater than end")
	}

	count := end - start
	for i := 0; i <= count; i++ {
		listFqdn = append(listFqdn, fmt.Sprintf("%s%d.%s", hostname, start+i, domain))
	}

	return listFqdn, nil
}
