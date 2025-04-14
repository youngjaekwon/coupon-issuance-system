package stock

import (
	"context"
	"time"
)

func (s *service) PreWarmStock(ctx context.Context, start, end time.Time) error {
	campaigns, err := s.campaignRepository.FindStartingBetween(ctx, start, end)
	if err != nil {
		return err
	}

	for _, campaign := range campaigns {
		isPreWarmed, err := s.repository.IsStockPreWarm(ctx, campaign.ID.String())
		if err != nil {
			return err
		}
		if isPreWarmed {
			continue
		}
		err = s.repository.PreWarmStock(ctx, campaign.ID.String(), campaign.TotalCount)
		if err != nil {
			return err
		}
	}

	return nil
}
