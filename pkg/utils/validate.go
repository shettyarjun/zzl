package utils

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// IsAlive takes a URL, sends an HTTP GET request, and returns the HTTP status code
func IsAlive(wg *sync.WaitGroup, url string, timeout time.Duration) (error) {
    defer wg.Done()

    client := &http.Client{
        Timeout: timeout * time.Second, // Set a timeout for the request
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
        },    
    }

    // Send an HTTP GET request
    resp, err := client.Get(url)
    if err != nil {
        return err // Return 0 and the error if the request fails
    }
    defer resp.Body.Close()

    // Print out the results
    fmt.Printf("%s, %d\n", url, resp.StatusCode)

    // Return the HTTP status code
    return nil
}

// IsAlive takes a URL, sends an HTTP GET request, and returns the HTTP status code
// without async
func IsAliveNormal(url string, timeout time.Duration) (error) {
    client := &http.Client{
        Timeout: timeout * time.Second, // Set a timeout for the request
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
        },    
    }

    // Send an HTTP GET request
    resp, err := client.Get(url)
    if err != nil {
        return err // Return 0 and the error if the request fails
    }
    defer resp.Body.Close()

    // Print out the results
    fmt.Printf("%s, %d\n", url, resp.StatusCode)

    // Return the HTTP status code
    return nil
}

// RemoveWildcard checks if the input string starts with "*." and removes it if it does
func RemoveWildcard(input string) string {
    // Check if the string starts with "*."
    if strings.HasPrefix(input, "*.") {
        return strings.TrimPrefix(input, "*.") // Remove "*."
    }
    return input // Return the original string if no wildcard
}

// RemoveWildcardAndDuplicates processes an array of strings to remove wildcards and duplicates
func RemoveWildcardAndDuplicates(input []string) []string {
    seen := make(map[string]struct{}) // Use a map to track unique entries
    result := []string{}               // Slice to store the cleaned results

    for _, str := range input {
        // Remove wildcard prefix if it exists
        cleaned := RemoveWildcard(str)
        // Check for uniqueness
        if _, exists := seen[cleaned]; !exists {
            seen[cleaned] = struct{}{} // Mark this cleaned string as seen
            result = append(result, cleaned) // Add to result if it's unique
        }
    }

    return result // Return the array with cleaned and unique entries
}