package main

import (
	"fmt"
	"sync"
)

// Exercise: Web Crawler
// In this exercise you'll use Go's concurrency features to parallelize a web crawler.
//
// Modify the `Crawl` function to fetch URLs in parallel without fetching the same URL twice.
//
// Hint: you can keep a cache of the URLs that have been fetched on a map,
// but maps alone are not safe for concurrent use!

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher) {
// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:
// 	if depth <= 0 {
// 		return
// 	}
// 	body, urls, err := fetcher.Fetch(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("found: %s %q\n", url, body)
// 	for _, u := range urls {
// 		Crawl(u, depth-1, fetcher)
// 	}
// 	return
// }
//
// func main() {
// 	Crawl("https://golang.org/", 4, fetcher)
// }
// - - -

// - - - Fetch URLs in parallel: start - - -
// 实现参考地址(https://stackoverflow.com/questions/12224962/exercise-web-crawler-concurrency-not-working)
// Also, don't rely on the map being thread safe.
type fetchedUrls struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, used *fetchedUrls, wg *sync.WaitGroup) {
	if depth <= 0 {
		return
	}

	used.mux.Lock()
	if used.urls[url] == false {
		used.urls[url] = true
		wg.Add(1)
		go func() {
			defer wg.Done()
			body, urls, err := fetcher.Fetch(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("found: %s %q\n", url, body)
			for _, u := range urls {
				Crawl(u, depth-1, fetcher, used, wg)
			}
			return
		}()
	}
	used.mux.Unlock()
	return
}

func main() {
	wg := &sync.WaitGroup{} // 官网(https://golang.org/pkg/sync/#WaitGroup)
	used := fetchedUrls{urls: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &used, wg)
	wg.Wait()
}

// - - - Fetch URLs in parallel: end - - -

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
