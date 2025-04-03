package report

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Report struct {
	URL string
}

const DefaultURL = "https://www.semace.ce.gov.br/boletim-de-balneabilidade"

func SeaWater(url string) (Report, error) {
	if url == "" {
		url = DefaultURL
	}

	c := colly.NewCollector()
	var reports []Report

	c.OnHTML("ul[class='ListaEst -Verde'] a", func(e *colly.HTMLElement) {
		reports = append(reports, Report{URL: e.Attr("href")})
	})

	c.Visit(url)

	if len(reports) == 0 {
		return Report{}, fmt.Errorf("Failed to find the last SEMACE's report")
	}

	return reports[1], nil
}
