package grpccommon

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	TraceIDLogKey = "trace_id"
)

func AttachTraceIDInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, // Must be gctx.Ctx
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		traceID := gtrace.GetTraceID(ctx)
		md := metadata.Pairs(TraceIDLogKey, traceID)
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func ExtractTraceIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var traceID string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if values := md.Get(TraceIDLogKey); len(values) > 0 {
				traceID = values[0]
			}
		}

		if traceID != "" {
			ctx, _ = gtrace.WithTraceID(ctx, traceID)
		}

		return handler(ctx, req)
	}
}
