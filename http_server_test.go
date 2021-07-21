package main

import (
	"net/http"
	"testing"
)

func TestServerReplies(t *testing.T) {
	t.Run("Server responds to GET", func(t *testing.T) {
		HttpServer()
		request, err := http.NewRequest(http.MethodGet, "localhost:3000", nil)

		if err != nil {
			t.Errorf("Shit's fucked %v", err)
		}

		got := request.Body

		print(got)

	})
}
