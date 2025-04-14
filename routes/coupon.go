package routes

import (
	"couponIssuanceSystem/gen/coupon/v1/couponv1connect"
	handler "couponIssuanceSystem/internal/handler/coupon"
	"couponIssuanceSystem/internal/service/coupon"
	"github.com/gin-gonic/gin"
)

func RegisterCouponRoutes(r *gin.Engine, service coupon.Service) {
	couponHandler := handler.NewHandler(service)
	path, couponServiceHandler := couponv1connect.NewCouponServiceHandler(couponHandler)
	r.Any(path+"*any", gin.WrapH(couponServiceHandler))
}
