package ssl 

import (
    "crypto/tls"
    "fmt"
    "net"
    "time"
)

// SANs retrieves and validates the SSL certificate of the given domain
// Works for both domains and IP addresses don't worry
func SANs(domain string) ([]string, error) {
    // Set up a custom dialer with a timeout
	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}

    // Set up a TLS connection to the specified domain on port 443
    conn, err := tls.DialWithDialer(dialer ,"tcp", fmt.Sprintf("%s:443", domain), &tls.Config{
        InsecureSkipVerify: true, // Skip verification; we only want to fetch the certificate
    })
    if err != nil {
        return nil, fmt.Errorf("failed to connect to %s: %v", domain, err)
    }
    defer conn.Close()

    // Get the first certificate in the chain
    cert := conn.ConnectionState().PeerCertificates[0]

    // Extract the Subject Alternative Names (SANs)
    sans := cert.DNSNames
    return sans, nil
}