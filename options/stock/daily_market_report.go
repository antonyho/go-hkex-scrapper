package stock

import "time"

// DailyMarketReport for Stock Options from HKEX
type DailyMarketReport struct {
	ContractMonth        string
	Price                float64
	Type                 string // C or P: Calls or Puts
	OpeningPrice         float64
	DailyHigh            float64
	DailyLow             float64
	SettlementPrice      float64
	ChangeInSettlement   float64
	IVPercent            uint
	Volume               uint
	OpenInterest         uint
	ChangeInOpenInterest int
}

func (r *DailyMarketReport) ContractMonthTime() (time.Time, error) {
	return time.Parse("Jan06", r.ContractMonth)
}
