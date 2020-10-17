package test

import (
	logic "../go/logic"
	model "../go/model"
	"testing"
	"time"
)

func TestLogic_fundOfMoreThan5k_shouldFail(t *testing.T) {
	// arrange
	fund1 := model.Fund{FundId: 1, CustomerId: 1, LoadAmount: 5001, Time: time.Now()}
	// act
	res, _ := logic.FundAccount(fund1)
	// assert
	if res {
		t.Errorf("Expected a fund of 5k to fail")
	}
}

func TestLogic_4fundsAday_shouldFail(t *testing.T) {
	// arrange
	logic.FundAccount(model.Fund{FundId: 1, CustomerId: 1, LoadAmount: 1, Time: time.Now()})
	logic.FundAccount(model.Fund{FundId: 1, CustomerId: 1, LoadAmount: 1, Time: time.Now()})
	logic.FundAccount(model.Fund{FundId: 1, CustomerId: 1, LoadAmount: 1, Time: time.Now()})
	// act
	res, _ := logic.FundAccount(model.Fund{FundId: 1, CustomerId: 1, LoadAmount: 1, Time: time.Now()})
	// assert
	if res {
		t.Errorf("Expected 4 funds a day to fail")
	}
}

func TestLogic_3fundsEveryDay_shouldPass(t *testing.T) {
	// arrange
	for i := 1; i < 5; i++ {
		duration := time.Duration(i) * 24 * time.Hour
		for j := 0; j < 3; j++ {
			// act
			res, _ := logic.FundAccount(model.Fund{
				FundId:     i*31 + j,
				CustomerId: 1,
				LoadAmount: 1,
				Time:       time.Now().Add(duration)})
			if !res {
				t.Errorf("Expected 3 funds a day to pass")
			}
		}
	}
}
