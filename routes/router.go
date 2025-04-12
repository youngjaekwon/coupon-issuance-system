package routes

import (
	"couponIssuanceSystem/gen/health/v1/healthv1connect"
	"couponIssuanceSystem/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := &handler.HealthServer{}
	path, healthServiceHandler := healthv1connect.NewHealthServiceHandler(healthHandler)

	r.Any(path+"*any", gin.WrapH(healthServiceHandler))

	return r
}
