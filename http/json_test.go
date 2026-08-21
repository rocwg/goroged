package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gedhttp "github.com/rocwg/ged/http"
)

type testJSONRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDecodeJSON(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		strings.NewReader(`{"name":"roc","age":36}`),
	)

	result, err := gedhttp.DecodeJSON[testJSONRequest](req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Name != "roc" {
		t.Fatalf(
			"got name %q, want %q",
			result.Name,
			"roc",
		)
	}

	if result.Age != 36 {
		t.Fatalf(
			"got age %d, want %d",
			result.Age,
			36,
		)
	}
}

func TestDecodeJSON_InvalidJSON(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		strings.NewReader(`{"name":`),
	)

	_, err := gedhttp.DecodeJSON[testJSONRequest](req)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDecodeJSON_EmptyBody(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		nil,
	)

	_, err := gedhttp.DecodeJSON[testJSONRequest](req)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
