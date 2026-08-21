package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gedhttp "github.com/rocwg/ged/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestWriteError(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{
			name:       "invalid argument",
			err:        status.Error(codes.InvalidArgument, "invalid request"),
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "unauthenticated",
			err:        status.Error(codes.Unauthenticated, "unauthenticated"),
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "permission denied",
			err:        status.Error(codes.PermissionDenied, "permission denied"),
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "not found",
			err:        status.Error(codes.NotFound, "not found"),
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "unavailable",
			err:        status.Error(codes.Unavailable, "provider unavailable"),
			wantStatus: http.StatusBadGateway,
		},
		{
			name:       "unknown",
			err:        errors.New("unexpected error"),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			gedhttp.WriteError(
				recorder,
				tt.err,
			)

			if recorder.Code != tt.wantStatus {
				t.Fatalf(
					"got status %d, want %d",
					recorder.Code,
					tt.wantStatus,
				)
			}
		})
	}
}

func TestWriteError_ResponseBody(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := status.Error(
		codes.NotFound,
		"area not found",
	)

	gedhttp.WriteError(
		recorder,
		err,
	)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusNotFound,
		)
	}

	body := recorder.Body.String()

	want := `"code":"NotFound"`
	if !strings.Contains(body, want) {
		t.Fatalf(
			"response body %q does not contain %q",
			body,
			want,
		)
	}

	want = `"message":"area not found"`
	if !strings.Contains(body, want) {
		t.Fatalf(
			"response body %q does not contain %q",
			body,
			want,
		)
	}
}
