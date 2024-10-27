package args

import (
    "flag"
)

// Config holds the parsed arguments
type Config struct {
    Timeout 	int
    Delay    	int
	Match 		string
	Output 		string
    Domain      string
    OnlyHttps   bool
    OnlyHttp    bool
    StartIp     string
    EndIp       string
}

// ParseArgs parses the command-line arguments using the flag library.
func ParseArgs() (*Config, error) {
    timeout := flag.Int("timeout", 1, "Max seconds to wait until we get a response from the server")
    delay := flag.Int("delay", 0, "Delay between each request, in seconds")
	match := flag.String("match", "", "Word to match the results with, only matched results are printed")
	output := flag.String("output", "", "Output file to save the results to")
    domain := flag.String("domain", "", "Subdoamin/domain to get the SANs from")
    onlyHttps := flag.Bool("only-https", false, "Use HTTPs protocol only while validating SANs")
    onlyHttp := flag.Bool("only-http", false, "Use HTTPs protocol only while validating SANs")
	startIp := flag.String("start-ip", "", "IP range to start with")
	endIp := flag.String("end-ip", "", "IP range to end with")

    // Parse the flags
    flag.Parse()

    // Create Config with parsed flags
    config := &Config{
        Timeout: 	*timeout,
        Delay:	    *delay,
		Match: 		*match,
		Output: 	*output,
        Domain:     *domain,
        OnlyHttps:  *onlyHttps,
        OnlyHttp:   *onlyHttp,
		StartIp: 	*startIp,
        EndIp:      *endIp,
    }

    return config, nil
}
