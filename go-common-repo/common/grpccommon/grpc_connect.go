package grpccommon

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DefaultGrpcDialOption    = grpc.WithTransportCredentials(insecure.NewCredentials())
	DefaultGrpcServiceConfig = `{"loadBalancingConfig": [{"round_robin":{}}]}`
	DefaultDialTimeout       = 5 * time.Second
)

func MustNewGrpcClientConn(ctx context.Context, addr string, opts ...grpc.DialOption) *grpc.ClientConn {
	if len(opts) == 0 {
		opts = []grpc.DialOption{DefaultGrpcDialOption}
	}

	opts = append(opts, grpc.WithDefaultServiceConfig(DefaultGrpcServiceConfig))

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	return conn
}

func NewGrpcClientConn(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	hasTransport := false

	for _, opt := range opts {
		if _, ok := opt.(interface{ IsTransportCredentials() }); ok {
			hasTransport = true
			break
		}
	}

	finalOpts := []grpc.DialOption{
		grpc.WithDefaultServiceConfig(DefaultGrpcServiceConfig),
		grpc.WithUnaryInterceptor(AttachTraceIDInterceptor()),
	}

	if !hasTransport {
		finalOpts = append(finalOpts, DefaultGrpcDialOption)
	}

	finalOpts = append(finalOpts, opts...)

	conn, err := grpc.DialContext(ctx, addr, finalOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
