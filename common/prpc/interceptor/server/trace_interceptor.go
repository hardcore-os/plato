package server

import (
	"context"

	ptrace "github.com/hardcore-os/plato/common/prpc/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// TraceUnaryServerInterceptor ...
func TraceUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md := metadata.MD{}
		header, ok := metadata.FromIncomingContext(ctx)

		if ok {
			md = header.Copy()
		}

		spanCtx := ptrace.Extract(ctx, otel.GetTextMapPropagator(), &md)
		tr := otel.Tracer(ptrace.TraceName)
		name, attrs := ptrace.BuildSpan(info.FullMethod, ptrace.PeerFromCtx(ctx))

		ctx, span := tr.Start(trace.ContextWithRemoteSpanContext(ctx, spanCtx), name, trace.WithSpanKind(trace.SpanKindServer), trace.WithAttributes(attrs...))
		defer span.End()

		resp, err = handler(ctx, req)
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				span.SetStatus(codes.Error, s.Message())
				span.SetAttributes(ptrace.StatusCodeAttr(s.Code()))
			} else {
				span.SetStatus(codes.Error, err.Error())
			}
			return nil, err
		}

		return resp, nil
	}
}
