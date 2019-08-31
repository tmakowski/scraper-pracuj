package scrapers

import (
	"github.com/gocolly/colly"
	"strings"
)

// Used for adding all handlers for pracuj.pl
func (c *CollectorsClient) AddHandlersPracuj() {
	c.addPagingOffers()
	c.addOffersPracuj()
}

// HTML element handlers to go through next pages of offers and pass them to
// the offers' collector.
func (c *CollectorsClient) addPagingOffers() {
	c.cPage.OnHTML("div[class='pagination'] a[href]",
		func(el *colly.HTMLElement) {
			link := el .Attr("href")
			el.Request.Visit(link)
		})

	c.cPage.OnHTML("a[class$='title-link']",
		func(el *colly.HTMLElement) {
			link := el.Request.AbsoluteURL(el.Attr("href"))
			c.cOffer.Visit(link)
		})
}

// HTML element handler for reading the offer's description.
func (c *CollectorsClient) addOffersPracuj() {
	c.cOffer.OnHTML("table[id='offTable']",
		func(el *colly.HTMLElement) {
			for _, k := range c.Keywords {
				if strings.Contains(strings.ToLower(el.Text), strings.ToLower(k)) {
					c.Data[k] = append(c.Data[k], el.Request.URL.String())
				}
			}
			c.OffersCounter++
		})
}
