package handler

import (
	"connectrpc.com/connect"
	"context"
	healthv1 "couponIssuanceSystem/gen/health/v1"
)

type HealthServer struct{}

func (s *HealthServer) Ping(
	ctx context.Context,
	req *connect.Request[healthv1.PingRequest],
) (*connect.Response[healthv1.PingResponse], error) {
	res := &healthv1.PingResponse{Message: "pong"}
	return connect.NewResponse(res), nil
}
