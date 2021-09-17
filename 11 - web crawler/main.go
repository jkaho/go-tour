package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	mu     sync.Mutex
	record map[string]bool // record of crawled sites
}

// Check whether site has already been crawled
func (c *Crawler) check(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.record[url]; ok {
		return true
	}
	c.record[url] = true

	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	crawled := c.check(url)
	if crawled {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err) // not found
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			c.Crawl(url, depth-1, fetcher)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

func main() {
	c := Crawler{
		record: make(map[string]bool),
	}
	c.Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
