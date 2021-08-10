package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerReplies(t *testing.T) {
	t.Run("Server responds to GET", func(t *testing.T) {
		_, err := http.NewRequest(http.MethodGet, "localhost:3000/", nil)
		response := httptest.NewRecorder()

		if err != nil {
			t.Errorf("Shit broke: %v", err)
		}

		got := response.Body.String()
		want := "200"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
