package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	edgehttp "github.com/rocwg/ged/http"
)

type testRequest struct {
	Name string `json:"name"`
}

type testResponse struct {
	Message string `json:"message"`
}

func TestJSONHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.POST(
		"/hello",
		edgehttp.JSONHandler[
			testRequest,
			testResponse,
		](
			func(
				c *gin.Context,
				req *testRequest,
			) error {

				return c.ShouldBindJSON(req)
			},
			func(
				ctx context.Context,
				req *testRequest,
			) (*testResponse, error) {

				return &testResponse{
					Message: "hello " + req.Name,
				}, nil
			},
		),
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/hello",
		strings.NewReader(`{"name":"roc"}`),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusOK,
		)
	}
}

func TestJSONHandler_BadRequest(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.POST(
		"/hello",
		edgehttp.JSONHandler[
			testRequest,
			testResponse,
		](
			func(
				c *gin.Context,
				req *testRequest,
			) error {

				return c.ShouldBindJSON(req)
			},
			func(
				ctx context.Context,
				req *testRequest,
			) (*testResponse, error) {

				return &testResponse{
					Message: "hello",
				}, nil
			},
		),
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/hello",
		nil, //strings.NewReader(`{"name":"roc"}`),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusBadRequest,
		)
	}
}
