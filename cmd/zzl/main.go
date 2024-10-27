package main

import (
	"fmt"
	"sync"
	"time"
	"log"

	"github.com/DEMON1A/zzl/pkg/args"
	"github.com/DEMON1A/zzl/pkg/ssl"
	"github.com/DEMON1A/zzl/pkg/utils"
	"github.com/DEMON1A/zzl/pkg/ip"
)

func main() {
	config, err := args.ParseArgs()
    if err != nil {
        fmt.Println("Error parsing arguments:", err)
        return
    }

	if config.Domain != "" {
		// collect sans 
		sanz, err := ssl.SANs(config.Domain)
		if err != nil {
			fmt.Println("Error getting DNS names:", err)
			return
		}

		// clean sans
		zanz := utils.RemoveWildcardAndDuplicates(sanz)

		// setup the wait group for concurrency
		var wg sync.WaitGroup

		for _, san := range zanz {
			if config.OnlyHttps {
				wg.Add(1) // add a single task
				go utils.IsAlive(&wg, fmt.Sprintf("https://%s", san), time.Duration(config.Timeout))
			} else if config.OnlyHttp {
				wg.Add(1) // add a single task
				go utils.IsAlive(&wg, fmt.Sprintf("http://%s", san), time.Duration(config.Timeout))
			} else {
				httpURL := fmt.Sprintf("http://%s", san)
				httpsURL := fmt.Sprintf("https://%s", san)

				wg.Add(2) // Add two tasks to the wait group
				go utils.IsAlive(&wg, httpURL, time.Duration(config.Timeout))
				go utils.IsAlive(&wg, httpsURL, time.Duration(config.Timeout))
			}
		}

		wg.Wait() // Wait for all checks to complete
	} else if config.StartIp != "" {
		// Assuming config struct is set up
		endIp, err := ip.GenerateEndIP(config.StartIp)
		if err != nil {
			log.Fatal(err)
		}

		// Generate IP addresses
		addresses, err := ip.GenerateIPs(config.StartIp, endIp)
		if err != nil {
			log.Fatal(err)
		}

		// Task channel for distributing tasks
		taskChan := make(chan string, len(addresses)*2) // Buffered for better performance

		// Set up wait group and start a fixed number of worker goroutines
		var wg sync.WaitGroup
		workerCount := 10 // Adjust based on your system’s capacity
		for i := 0; i < workerCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for url := range taskChan {
					// Perform check and timeout
					utils.IsAliveNormal(url, time.Duration(config.Timeout))
				}
			}()
		}

		// Distribute tasks
		for _, address := range addresses {
			sanz, err := ssl.SANs(address)
			if err != nil {
				continue // Skip to the next address if error
			}

			// Clean up SAN entries
			zanz := utils.RemoveWildcardAndDuplicates(sanz)

			// Add each unique SAN to task channel with appropriate protocol
			for _, san := range zanz {
				if config.OnlyHttps {
					taskChan <- fmt.Sprintf("https://%s", san)
				} else if config.OnlyHttp {
					taskChan <- fmt.Sprintf("http://%s", san)
				} else {
					taskChan <- fmt.Sprintf("http://%s", san)
					taskChan <- fmt.Sprintf("https://%s", san)
				}
			}
		}

		// Close task channel once all tasks are added
		close(taskChan)

		// Wait for all workers to finish
		wg.Wait()
	} else {
		log.Fatal("You must either provide an IP range or a domain")
	}
}