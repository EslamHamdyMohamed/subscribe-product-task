package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"os"
	"strconv"
	"subscribe-product/app/exchange"
	"subscribe-product/app/models"
	"subscribe-product/app/services"
)

func Jobs(s *gocron.Scheduler) {
	s.Every(10).Second().Do(generateRequiredBills)
}

func generateRequiredBills() {
	limit, _ := strconv.Atoi(os.Getenv("CRON_SUBSCRIPTIONS_LIMIT"))
	offset := 0
	for {
		_, subscriptions := models.GetRequiredSubscriptions(exchange.PAID, limit, offset)
		fmt.Println(">>>>WELCOME: ", len(subscriptions))
		if len(subscriptions) == 0 {
			break
		}
		for _, subscription := range subscriptions {
			request := SetSubscriptionRequest(subscription)
			services.GenerateBill(request, exchange.PENDING)
			services.UpdateBill(subscription, exchange.EXPIRED)
		}
		offset++
	}
}
