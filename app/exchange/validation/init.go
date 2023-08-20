package validation

import appHelper "subscribe-product/app/helpers"

var TimeIntervalsMap = map[appHelper.TimeInterval]bool{}

func init() {
	TimeIntervalsMap = map[appHelper.TimeInterval]bool{
		appHelper.MONTHLY: true,
		appHelper.YEARLY:  true,
	}
}
