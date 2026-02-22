package scraper

import (
	"github.com/gocolly/colly"
)

type RawCardData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Attribute string `json:"attribute"`
	Level     string `json:"level"`
	Attack    string `json:"attack"`
	Defense   string `json:"defense"`
	Rarity    string `json:"rarity"`
	Text      string `json:"text"`
}

func parseCardData(idx int, e *colly.HTMLElement, data *RawCardData) {
	switch idx {
	case 0:
		data.ID = e.Text
	case 1:
		data.Name = e.Text
	case 2:
		data.Type = e.Text
	case 3:
		data.Attribute = e.Text
	case 4:
		data.Subtype = e.Text
	case 5:
		data.Level = e.Text
	case 6:
		data.Attack = e.Text
	case 7:
		data.Defense = e.Text
	case 8:
		data.Rarity = e.Text
	case 9:
		data.Text = e.Text
	}
}

func Visit(url string) ([]RawCardData, error) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) ... Safari/537.36")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,...")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Connection", "keep-alive")
	})

	cards := []RawCardData{}

	c.OnHTML("table#card_data", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			data := RawCardData{}
			h.ForEach("td", func(i int, g *colly.HTMLElement) {
				parseCardData(i, g, &data)
			})
			cards = append(cards, data)
		})
	})
	return cards, c.Visit(url)
}
