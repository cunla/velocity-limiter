package logic

import (
	model "../model"
	"fmt"
	"time"
)

var fundModel model.FundModel

const durationDay = -24 * time.Hour
const durationWeek = -7 * 24 * time.Hour

func startOfWeek(t time.Time) time.Time {
	var diff = time.Duration((6+int64(t.Weekday()))%7) * durationDay
	return t.Add(diff)
}

func FundAccount(fund model.Fund) (bool, error) {
	prev, err := fundModel.FindById(fund.CustomerId, fund.FundId)
	if err == nil {
		return prev.Accepted, fmt.Errorf("fund exists")
	}
	daySum, dayCount := fundModel.CountAndSumOfFundsOnDay(fund.CustomerId, fund.Time)
	weekStart := startOfWeek(fund.Time)
	weekSum := fundModel.SumOfFundsBetween(
		fund.CustomerId,
		weekStart,
		weekStart.Add(durationWeek))
	fund.Accepted = (dayCount < 3) && (daySum+fund.LoadAmount <= 5000.0) && (weekSum+fund.LoadAmount < 20000.0)
	fundModel.Create(&fund)
	return fund.Accepted, nil
}
