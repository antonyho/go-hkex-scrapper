package hsi

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDailyMarketReport_ContractMonthTime(t *testing.T) {
	r := DailyMarketReport{ContractMonth: "MAR-19"}
	month, err := r.ContractMonthTime()
	assert.NoError(t, err)
	assert.EqualValues(t, time.Date(2019, 3, 1, 0, 0, 0, 0, time.UTC), month)
}
