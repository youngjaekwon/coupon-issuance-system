package routes

import (
	"couponIssuanceSystem/internal/service/campaign"
	"couponIssuanceSystem/internal/service/coupon"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	campaignService campaign.Service,
	couponService coupon.Service,
) *gin.Engine {
	r := gin.Default()

	RegisterHealthRoutes(r)
	RegisterCampaignRoutes(r, campaignService)
	RegisterCouponRoutes(r, couponService)

	return r
}
