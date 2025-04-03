package report

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func startServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
			<html>
			<head>
			<title>Test Page</title>
			</head>
			<body>
			<div>
			<ul class="ListaEst -Verde">
			<li><a href="semace.first.url"></a></li>
			</ul>
			<ul class="ListaEst -Verde">
			<li><a href="semace.second.url"></a></li>
			</ul>
			</div>
			</body>
			</html>
			`))
	})

	return httptest.NewServer(mux)
}

func TestReportURL(t *testing.T) {
	srv := startServer()
	defer srv.Close()

	r, _ := SeaWater(srv.URL)

	if !strings.Contains(r.URL, "semace.second.url") {
		t.Errorf("report.URL = %q, want %q", r.URL, "semace.second.url")
	}
}

func TestReportURLNotFound(t *testing.T) {
	r, _ := SeaWater("https://blablalbla.ble")

	if r.URL != "" {
		t.Errorf("report.URL = %q, want %q", r.URL, "")
	}
}
