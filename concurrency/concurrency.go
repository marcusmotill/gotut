package concurrency

import (
	"fmt"
	"sync"
)

//Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, mux *sync.Mutex) {
	var err error
	var urls []string

	if depth <= 0 {
		return
	}

	mux.Lock()

	foundUrls[url]++
	if foundUrls[url] == 1 {
		_, urls, err = fetcher.Fetch(url)
	}

	mux.Unlock()

	if err != nil {
		return
	}

	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher, mux)
			done <- true
		}(u)
	}

	//wait for the urls to finish
	for range urls {
		<-done
	}

	return
}

// Run exported
func Run() {
	mux := &sync.Mutex{}

	Crawl("https://golang.org/", 4, fetcher, mux)
	for url, count := range foundUrls {
		fmt.Printf("found: %s %d\n", url, count)
	}
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

var foundUrls = make(map[string]int)

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
