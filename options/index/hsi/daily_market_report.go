package hsi

import "time"

// DailyMarketReport for HSI Options from HKEX
type DailyMarketReport struct {
	ContractMonth string
	StrikePrice   uint
	Type          string // C or P (Call or Put)
	// After-Hours Trading Session
	AfterHoursOpeningPrice uint
	AfterHoursDailyHigh    uint
	AfterHoursDailyLow     uint
	AfterHoursClosePrice   uint
	AfterHoursVolume       uint
	// Day Trading Session
	DayOpeningPrice uint
	DayDailyHigh    uint
	DayDailyLow     uint
	DayOQPClose     uint
	DayOQPChange    int
	DayIVPercent    int
	DayVolume       uint
	// Combined
	ContractHigh         uint
	ContractLow          uint
	Volume               uint
	OpenInterest         uint
	ChangeInOpenInterest int
}

func (r *DailyMarketReport) ContractMonthTime() (time.Time, error) {
	return time.Parse("Jan-06", r.ContractMonth)
}
