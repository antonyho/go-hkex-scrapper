package stock

import "time"

// DailyMarketReport for Stock Options from HKEX
type DailyMarketReport struct {
	ContractMonth string
}

func (r *DailyMarketReport) ContractMonthTime() (time.Time, error) {
	return time.Parse("Jan06", r.ContractMonth)
}
