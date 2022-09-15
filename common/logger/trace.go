package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

const (
	traceID = "trace_id"
)

func GetTraceID(ctx context.Context) string {
	var traceID string
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().TraceID().IsValid() {
		traceID = span.SpanContext().TraceID().String()
	}

	return traceID
}
