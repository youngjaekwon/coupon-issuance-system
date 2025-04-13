package coupon

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, coupon *models.Coupon) error
	ExistsByUser(ctx context.Context, campaignID uuid.UUID, userID string) (bool, error)
}
