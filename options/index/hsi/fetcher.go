package hsi

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/antonyho/go-hkex-scrapper/options/index"

	"github.com/antonyho/go-hkex-scrapper/options"
)

const (
	SourceURLPattern = "https://www.hkex.com.hk/eng/stat/dmstat/dayrpt/hsio%s.csv" // date string pattern YYMMDD
)

var (
	ErrNoReportAvailable = errors.New("no report available")
)

func Get(date time.Time) ([]index.DailyMarketReport, error) {
	resp, err := http.Get(fmt.Sprintf(SourceURLPattern, options.FormatURLDate(date)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, ErrNoReportAvailable
	}
	return index.ParseCSV(resp.Body)
}
