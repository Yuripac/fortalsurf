package main

import (
	"fmt"
	"fortalsurf/notifier"
	"io"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"log"
)

type Report struct {
	URL string
}

func FetchSemaceReport() (Report, error) {
	c := colly.NewCollector()
	var reports []Report

	c.OnHTML("ul[class='ListaEst -Verde'] a", func(e *colly.HTMLElement) {
		reports = append(reports, Report{URL: e.Attr("href")})
	})

	c.Visit("https://www.semace.ce.gov.br/boletim-de-balneabilidade")

	if len(reports) == 0 {
		return Report{}, fmt.Errorf("Failed to find the last SEMACE's report")
	}

	return reports[1], nil
}


func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	report, err := FetchSemaceReport()

	text := report.URL
	if err != nil {
		text = err.Error()
	}

	resp, err := notifier.Send(notifier.NewTelegram(), text)
	if err != nil {
		panic(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Notifier response:", string(respBody))
}

