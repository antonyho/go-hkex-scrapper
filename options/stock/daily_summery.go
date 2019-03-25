package stock

// DailySummary summarise the performance on each stock options
type DailySummary struct {
	Code               string
	Name               string
	Symbol             uint
	TradingVolumeTotal uint
	TradingVolumeCalls uint
	TradingVolumePuts  uint
	OpenInterestTotal  uint
	OpenInterestCalls  uint
	OpenInterestPuts   uint
	IVPercent          uint
}
