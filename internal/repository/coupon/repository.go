package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
)

type Repository interface {
	Create(ctx context.Context, coupon *models.Coupon) error
}
