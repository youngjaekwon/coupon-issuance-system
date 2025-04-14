package routes

import (
	"couponIssuanceSystem/internal/service/campaign"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	campaignService campaign.Service,
) *gin.Engine {
	r := gin.Default()

	RegisterHealthRoutes(r)
	RegisterCampaignRoutes(r, campaignService)

	return r
}
