package validation

import (
	appHelper "subscribe-product/app/exchange"
)

var TimeIntervalsMap = map[appHelper.TimeInterval]bool{}

func init() {
	TimeIntervalsMap = map[appHelper.TimeInterval]bool{
		appHelper.MONTHLY: true,
		appHelper.YEARLY:  true,
	}
}
