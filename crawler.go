package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

/*
Modelar o objeto soundtrack

Soundtrack: uma lista de musicas relacionadas com um movie/tvshow
- movie/tvshow
- soundtrack
- music

*/

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// c.AllowedDomains = []string{"what-song.com/tvshow/browse", "what-song.com/movies"}

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.what-song.com/movies/browse")
}
