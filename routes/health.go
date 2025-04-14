package routes

import (
	"couponIssuanceSystem/gen/health/v1/healthv1connect"
	"couponIssuanceSystem/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	healthHandler := &handler.HealthServer{}
	path, healthServiceHandler := healthv1connect.NewHealthServiceHandler(healthHandler)

	r.Any(path+"*any", gin.WrapH(healthServiceHandler))

	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
