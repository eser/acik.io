package grpcfx

import (
	"context"

	"github.com/eser/acik.io/pkg/bliss/metricsfx"
	"google.golang.org/grpc"
)

func MetricsInterceptor(mp metricsfx.MetricsProvider) grpc.UnaryServerInterceptor {
	// metrics := NewMetrics(mp)
	//
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// startTime := time.Now()
		//
		resp, err := handler(ctx, req)
		//
		// duration := time.Since(startTime)
		// st, _ := status.FromError(err)
		//
		// metrics.RequestsTotal.WithLabelValues(
		// 	info.FullMethod,
		// 	info.Server.String(),
		// 	st.Code().String(),
		// ).Inc()
		//
		// metrics.RequestDuration.WithLabelValues(
		// 	info.FullMethod,
		// 	info.Server.String(),
		// ).Observe(duration.Seconds())

		return resp, err
	}
}
