package main

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	urls := []string{
		"https://www.netflix.com/title/81260637/",
		"https://www.netflix.com/title/81056700/",
		"https://www.netflix.com/title/80172819/",
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "pt-BR,pt;q=0.9,en;q=0.8")
	})

	var wg sync.WaitGroup

	c.OnHTML(".title-title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("Title:", title)
	})

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			err := c.Visit(u)
			if err != nil {
				fmt.Printf("Error when accessing the URL %s: %v\n", u, err)
			}
		}(url)
	}

	wg.Wait()
}
