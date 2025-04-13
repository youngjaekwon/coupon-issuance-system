package stock

import "fmt"

func StockKey(campaignID string) string {
	return fmt.Sprintf("coupon:campaign:%s:stock", campaignID)
}
