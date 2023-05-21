package mdw

import (
	"context"
	"net/http"

	"github.com/segmentio/ksuid"
)

const (
	// RequestIDHeader is a header name for request ID.
	RequestIDHeader = "X-Request-ID"
)

// contextKey is a type for context key.
type contextKey struct{ name string }

// String returns context key as string.
func (c *contextKey) String() string {
	return "httpx_context_value_" + c.name
}

// RequestIDContextKey is a context key for request ID.
var RequestIDContextKey = &contextKey{name: "request_id"}

// RequestID is a middleware that adds request ID to the request context.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = ksuid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequestIDContextKey, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID returns request ID from the request context.
func GetRequestID(ctx context.Context) string {
	return ctx.Value(RequestIDContextKey).(string)
}
