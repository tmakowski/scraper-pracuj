package scrapers

import (
	"github.com/gocolly/colly"
)

// Client for collectors: one is used for crawling pages and the other is used
// for reading offers.
type CollectorsClient struct {
	cPage         *colly.Collector
	cOffer        *colly.Collector
	City          string
	Keywords      []string
	OffersCounter int
	Data          map[string][]string
}

// Function used to create the CollectorsClient.
func CreateClient(city string, keywords []string) *CollectorsClient {
	cPage := colly.NewCollector(
		colly.Async(true))
	cOffer := cPage.Clone()

	c := &CollectorsClient{cPage, cOffer, city, keywords, 0, map[string][]string{}}
	for _, k := range c.Keywords {
		c.Data[k] = []string{}
	}

	return c
}

// Visit wrapper for CollectorsClient. Starts the underlying pages' collector.
func (c *CollectorsClient) Visit(url string) error {
	return c.cPage.Visit(url)
}

// Wait wrapper for Collectors Client. Waits for both collectors to finish.
func (c *CollectorsClient) Wait() {
	c.cPage.Wait()
	c.cOffer.Wait()
}
