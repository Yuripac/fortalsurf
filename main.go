package main

import (
	"fmt"
	"fortalsurf/notifier"
	"io"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
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

func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	report, err := FetchSemaceReport()
	//
	// 	text := report.URL
	// 	if err != nil {
	// 		text = err.Error()
	// 	}
	//
	// 	resp, err := notifier.Send(notifier.NewTelegram(), text)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	respBody, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	log.Println("Notifier response:", string(respBody))
	// })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	srv := http.Server{
		Addr: ":"+port,
		Handler: http.HandlerFunc(handler),
		WriteTimeout:      30 * time.Second,
		ReadTimeout:       30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// log.Printf("listing on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}

