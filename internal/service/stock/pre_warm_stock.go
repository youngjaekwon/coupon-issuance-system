package stock

import (
	"context"
	"github.com/google/uuid"
)

func (s *service) PreWarmStock(ctx context.Context, campaignID uuid.UUID, totalCount int) error {
	campaignIDStr := campaignID.String()
	isPreWarm, err := s.repository.IsStockPreWarm(ctx, campaignIDStr)
	if err != nil {
		return err
	}

	if isPreWarm {
		return nil
	}

	err = s.repository.PreWarmStock(ctx, campaignIDStr, totalCount)
	if err != nil {
		return err
	}

	return nil
}
