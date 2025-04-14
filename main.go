package main

import (
	"couponIssuanceSystem/internal/config"
	stockcron "couponIssuanceSystem/internal/cron/stock"
	"couponIssuanceSystem/internal/infra/db"
	"couponIssuanceSystem/internal/infra/redis"
	campaignrepo "couponIssuanceSystem/internal/repository/campaign"
	couponrepo "couponIssuanceSystem/internal/repository/coupon"
	stockrepo "couponIssuanceSystem/internal/repository/stock"
	campaignsvc "couponIssuanceSystem/internal/service/campaign"
	couponsvc "couponIssuanceSystem/internal/service/coupon"
	stocksvc "couponIssuanceSystem/internal/service/stock"
	"couponIssuanceSystem/internal/utils/couponcode"
	"couponIssuanceSystem/routes"
	"log"
	"net/http"
)

func main() {
	config.Init()
	couponcode.Init()
	dbConnection := db.Init()
	redisClient := redis.Init()

	campaignRepository := campaignrepo.New(dbConnection)
	couponRepository := couponrepo.New(dbConnection)
	stockRepository := stockrepo.New(redisClient)
	codeGenerator := couponcode.NewGenerator()

	campaignService := campaignsvc.New(campaignRepository)
	couponService := couponsvc.New(couponRepository, campaignService, stockRepository, codeGenerator)
	stockService := stocksvc.New(stockRepository, campaignRepository)
	stockWarmer := stockcron.NewWarmer(stockService)

	go stockWarmer.Run()

	mux := routes.SetupMux(campaignService, couponService)
	server := &http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
