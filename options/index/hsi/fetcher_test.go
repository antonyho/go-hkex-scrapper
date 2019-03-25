package hsi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("TradeDay", func(t *testing.T) {
		report, err := Get(time.Date(2019, 3, 25, 0, 0, 0, 0, time.UTC))
		assert.NoError(t, err)
		assert.NotEmpty(t, report)
	})
	t.Run("NonTradeDay", func(t *testing.T) {
		report, err := Get(time.Date(2019, 3, 24, 0, 0, 0, 0, time.UTC))
		if assert.Error(t, err) {
			t.Logf("Error: %+v", err)
		}
		assert.Empty(t, report)
	})
}
