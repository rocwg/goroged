package http

import (
	"encoding/json"
	nethttp "net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteError(
	w nethttp.ResponseWriter,
	err error,
) {
	code := status.Code(err)

	statusCode := grpcCodeToHTTP(code)

	_ = WriteJSON(
		w,
		statusCode,
		ErrorResponse{
			Code:    code.String(),
			Message: status.Convert(err).Message(),
		},
	)
}

func WriteJSON(
	w nethttp.ResponseWriter,
	statusCode int,
	value any,
) error {

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(data)
	return err
}

// gRPC                    HTTP
// ────────────────────────────────
// InvalidArgument         400
// Unauthenticated         401
// PermissionDenied        403
// NotFound                404
// AlreadyExists           409
// FailedPrecondition      400
// DeadlineExceeded        504
// Unavailable             502
// 其他                    500
func grpcCodeToHTTP(code codes.Code) int {
	switch code {
	case codes.InvalidArgument:
		return nethttp.StatusBadRequest

	case codes.Unauthenticated:
		return nethttp.StatusUnauthorized

	case codes.PermissionDenied:
		return nethttp.StatusForbidden

	case codes.NotFound:
		return nethttp.StatusNotFound

	case codes.AlreadyExists:
		return nethttp.StatusConflict

	case codes.FailedPrecondition:
		return nethttp.StatusBadRequest

	case codes.DeadlineExceeded:
		return nethttp.StatusGatewayTimeout

	case codes.Unavailable:
		return nethttp.StatusBadGateway

	default:
		return nethttp.StatusInternalServerError
	}
}
