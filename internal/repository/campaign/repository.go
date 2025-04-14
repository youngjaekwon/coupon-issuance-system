package campaign

import (
	"context"
	"couponIssuanceSystem/internal/models"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	Create(ctx context.Context, campaign *models.Campaign) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Campaign, error)
	List(ctx context.Context, page, limit int) ([]*models.Campaign, error)
	FindStartingBetween(ctx context.Context, start, end time.Time) ([]*models.Campaign, error)
}
