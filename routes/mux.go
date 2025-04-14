package routes

import (
	"couponIssuanceSystem/gen/campaign/v1/campaignv1connect"
	"couponIssuanceSystem/gen/coupon/v1/couponv1connect"
	"couponIssuanceSystem/gen/health/v1/healthv1connect"
	"couponIssuanceSystem/internal/handler"
	campaignhandler "couponIssuanceSystem/internal/handler/campaign"
	couponhandler "couponIssuanceSystem/internal/handler/coupon"
	"couponIssuanceSystem/internal/service/campaign"
	"couponIssuanceSystem/internal/service/coupon"
	"net/http"
)

func SetupMux(
	campaignService campaign.Service,
	couponService coupon.Service,
) *http.ServeMux {
	mux := http.NewServeMux()

	campaignHandler := campaignhandler.NewHandler(campaignService)
	campaignPath, campaignServiceHandler := campaignv1connect.NewCampaignServiceHandler(campaignHandler)
	mux.Handle(campaignPath, campaignServiceHandler)

	couponHandler := couponhandler.NewHandler(couponService)
	couponPath, couponServiceHandler := couponv1connect.NewCouponServiceHandler(couponHandler)
	mux.Handle(couponPath, couponServiceHandler)

	healthHandler := &handler.HealthServer{}
	healthPath, healthServiceHandler := healthv1connect.NewHealthServiceHandler(healthHandler)
	mux.Handle(healthPath, healthServiceHandler)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	return mux
}
