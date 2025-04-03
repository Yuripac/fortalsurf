package main

import (
	"fortalsurf/notifier"
	"fortalsurf/report"
	"io"
	"net/http"
	"os"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		report, err := report.SeaWater("")

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
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	log.Printf("listing on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

