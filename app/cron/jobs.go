package cron

import (
	"github.com/go-co-op/gocron"
	"os"
	"strconv"
	appHelper "subscribe-product/app/helpers"
	"subscribe-product/app/models"
	"subscribe-product/app/services"
)

func Jobs(s *gocron.Scheduler) {
	s.Every(1).Day().Do(generateRequiredBills)
}

func generateRequiredBills() {
	limit, _ := strconv.Atoi(os.Getenv("CRON_SUBSCRIPTIONS_LIMIT"))
	offset := 1
	for {
		_, subscriptions := models.GetRequiredSubscriptions(appHelper.PAID, limit, offset)
		if len(subscriptions) == 0 {
			break
		}
		for _, subscription := range subscriptions {
			request := SetSubscriptionRequest(subscription)
			services.GenerateBill(request, appHelper.PENDING)
		}
		offset++
	}
}
