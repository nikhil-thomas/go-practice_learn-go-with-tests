package concurrency

// WebsiteChecker checks validity of a url
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks the calidity of a list of urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultsChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		res := <-resultsChan
		results[res.string] = res.bool
	}

	return results
}
