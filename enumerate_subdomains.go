package subzero

import "sync"

func EnumerateSubdomains(domain string, options *EnumerationOptions) <-chan *Result {
	results := make(chan *Result)
	go func(domain string, options *EnumerationOptions, results chan *Result) {
		defer close(results)
		wg := sync.WaitGroup{}
		for _, source := range options.Sources {
			wg.Add(1)
			// merge source results channels with go funcs
			go func(source Source, results chan *Result) {
				defer wg.Done()
				for result := range source.ProcessDomain(domain) {
					results <- result
				}
			}(source, results)
		}
		wg.Wait()
	}(domain, options, results)
	return results
}
