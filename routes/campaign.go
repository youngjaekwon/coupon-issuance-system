package routes

import (
	"couponIssuanceSystem/gen/campaign/v1/campaignv1connect"
	handler "couponIssuanceSystem/internal/handler/campaign"
	"couponIssuanceSystem/internal/service/campaign"
	"github.com/gin-gonic/gin"
)

func RegisterCampaignRoutes(r *gin.Engine, campaignService campaign.Service) {
	campaignHandler := handler.NewHandler(campaignService)
	path, campaignServiceHandler := campaignv1connect.NewCampaignServiceHandler(campaignHandler)

	r.Any(path+"*any", gin.WrapH(campaignServiceHandler))
}
