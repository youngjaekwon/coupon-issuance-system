package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	const totalRequests = 1000
	var wg sync.WaitGroup
	url := "http://localhost:8000/coupon.v1.CouponService/IssueCoupon"

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			body := []byte(fmt.Sprintf(`{"campaign_id":"00000000-0000-0000-0000-000000000001","user_id":"user-%d"}`, i))
			resp, err := http.Post(url, "application/json", bytes.NewReader(body))
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			resp.Body.Close()
		}(i)
	}
	wg.Wait()
}
