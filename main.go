package main

import (
	"couponIssuanceSystem/routes"
	"log"
)

func main() {
	r := routes.SetupRouter()

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("서버 실행 중 오류 발생 : %v", err)
	}
}
