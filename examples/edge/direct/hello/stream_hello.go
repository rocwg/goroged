package hello

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

// streamHello
//
// HTTP(SSE)
//
//	↓
//
// Bridge
//
//	↓
//
// gRPC Server Streaming
func (a *Adapter) streamHello(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req StreamHelloRequest

	// 绑定 Query 参数，例如：?name=roc
	req.Name = r.URL.Query().Get("name")

	ctx, cancel := context.WithTimeout(
		r.Context(),
		a.streamTimeout,
	)
	defer cancel()

	stream, err := a.client.StreamHello(
		ctx,
		&hellov1.StreamHelloRequest{
			Name: req.Name,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 告诉浏览器：这是一个 SSE 持续流。
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	// SSE Flush
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "stream unsupported", http.StatusInternalServerError)
		return
	}

	// 持续接收 gRPC Stream
	for {
		resp, err := stream.Recv()

		// 服务端正常结束
		if err == io.EOF {
			return
		}

		// Stream 异常
		if err != nil {
			if writeErr := writeSSE(w, "error", err.Error()); writeErr != nil {
				return
			}
			flusher.Flush()
			return
		}

		// gRPC Stream -> HTTP SSE
		if err := writeSSE(w, "message", resp); err != nil {
			return
		}
		flusher.Flush()
	}
}

func writeSSE(
	w http.ResponseWriter,
	event string,
	data any,
) error {

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(
		w,
		"event: %s\n"+
			"data: %s\n\n",
		event,
		payload,
	)

	return err
}
