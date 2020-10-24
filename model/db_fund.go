package model

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Fund struct {
	gorm.Model
	FundId     int
	CustomerId int
	LoadAmount float64
	Time       time.Time
	Accepted   bool
}

func initialize_db() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Fund{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

var fundStore = initialize_db()

type FundModel struct{}

func (fundModel FundModel) CountAndSumOfFundsOnDay(customerId int, time time.Time) (float64, int32) {
	var sum sql.NullFloat64
	var count int32
	row := fundStore.Table("funds").Where(
		"accepted = true and customer_id = ? and time >= ?",
		customerId, time).Select("sum(load_amount), count(*)").Row()
	err := row.Scan(&sum, &count)
	if err != nil || !sum.Valid {
		return 0, 0
	}
	return sum.Float64, count
}

func (fundModel FundModel) SumOfFundsBetween(customerId int, from time.Time, to time.Time) float64 {
	var sum sql.NullFloat64
	row := fundStore.Table("funds").Where(
		"accepted = true and customer_id = ? and time >= ? and time < ?",
		customerId, from, to).Select(
		"sum(load_amount)").Row()
	err := row.Scan(&sum)
	if err != nil || !sum.Valid {
		return 0
	}
	return sum.Float64
}

func (fundModel FundModel) FindById(customerId int, fundId int) (Fund, error) {
	var fund Fund
	err := fundStore.Where(&Fund{CustomerId: customerId, FundId: fundId}).First(&fund).Error
	if err != nil {
		return fund, err
	}
	return fund, nil
}

func (fundModel FundModel) Create(fund *Fund) {
	fundStore.Create(fund)
}
