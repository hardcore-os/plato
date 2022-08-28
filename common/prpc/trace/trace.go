package trace

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
)

const (
	TraceName = "plato-trace"
)

type metadataSupplier struct {
	metadata *metadata.MD
}

func (s *metadataSupplier) Get(key string) string {
	values := s.metadata.Get(key)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func (s *metadataSupplier) Set(key, value string) {
	s.metadata.Set(key, value)
}

func (s *metadataSupplier) Keys() []string {
	out := make([]string, 0, len(*s.metadata))
	for key := range *s.metadata {
		out = append(out, key)
	}

	return out
}

// Inject set cross-cutting concerns from the Context into the metadata.
func Inject(ctx context.Context, p propagation.TextMapPropagator, m *metadata.MD) {
	p.Inject(ctx, &metadataSupplier{
		metadata: m,
	})
}

// Extract reads cross-cutting concerns from the metadata into a Context.
func Extract(ctx context.Context, p propagation.TextMapPropagator, metadata *metadata.MD) sdktrace.SpanContext {
	ctx = p.Extract(ctx, &metadataSupplier{
		metadata: metadata,
	})

	return sdktrace.SpanContextFromContext(ctx)
}
