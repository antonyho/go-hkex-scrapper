package stock

import "errors"

const (
	SourceURLPattern = "https://www.hkex.com.hk/eng/stat/dmstat/dayrpt/dqe%s.csv" // date string pattern YYMMDD
)

var (
	ErrNoReportAvailable = errors.New("no report available")
)
