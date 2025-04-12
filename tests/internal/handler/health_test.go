package handler__test

import (
	"connectrpc.com/connect"
	"context"
	healthv1 "couponIssuanceSystem/gen/health/v1"
	"couponIssuanceSystem/internal/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthServer_Ping(t *testing.T) {
	// Given
	s := &handler.HealthServer{}

	// When
	req := connect.NewRequest(&healthv1.PingRequest{})
	res, err := s.Ping(context.Background(), req)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "pong", res.Msg.Message)
}
