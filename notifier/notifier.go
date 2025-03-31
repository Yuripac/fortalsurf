package notifier

import "net/http"

type Notifier interface {
	Send(string) (*http.Response, error)
}

func Send(n Notifier, text string) (*http.Response, error) {
	return n.Send(text)
}

