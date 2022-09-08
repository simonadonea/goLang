package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites1(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result, 1)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// close(resultChannel)
	wait := 3
	for wait > 0 {
		r := <-resultChannel
		results[r.string] = r.bool
		wait--
	}

	return results
}

func CheckWebsites2(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
