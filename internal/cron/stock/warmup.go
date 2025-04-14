package stock

import (
	"context"
	"couponIssuanceSystem/internal/service/stock"
	"log"
	"time"
)

type Warmer struct {
	service stock.Service
}

func NewWarmer(service stock.Service) *Warmer {
	return &Warmer{
		service: service,
	}
}

func (w *Warmer) Run() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.preWarmStock()
		}
	}
}

func (w *Warmer) preWarmStock() {
	ctx := context.Background()

	now := time.Now()
	start := now.Add(-time.Minute * 5)
	end := now.Add(time.Minute * 5)

	err := w.service.PreWarmStock(ctx, start, end)
	if err != nil {
		log.Printf("[Prewarm] 실패: %v\n", err)
		return
	}
}
