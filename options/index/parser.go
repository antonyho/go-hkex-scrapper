package index

import (
	"encoding/csv"
	"io"
	"strconv"
)

func ParseCSV(reader io.Reader) ([]DailyMarketReport, error) {
	csvRdr := csv.NewReader(reader)
	csvRdr.Comment = '"'
	csvRdr.FieldsPerRecord = -1 // do not check field count for each row, the file is not regularly formatted
	rows, err := csvRdr.ReadAll()
	if err != nil {
		return nil, err
	}
	dailyMarketReport := []DailyMarketReport{}
	for _, row := range rows {
		if len(row) == 20 {
			strikePrice, _ := strconv.Atoi(row[1])
			afterHoursOpeningPrice, _ := strconv.Atoi(row[3])
			afterHoursDailyHigh, _ := strconv.Atoi(row[4])
			afterHoursDailyLow, _ := strconv.Atoi(row[5])
			afterHoursClosePrice, _ := strconv.Atoi(row[6])
			afterHoursVolume, _ := strconv.Atoi(row[7])
			dayOpeningPrice, _ := strconv.Atoi(row[8])
			dayDailyHigh, _ := strconv.Atoi(row[9])
			dayDailyLow, _ := strconv.Atoi(row[10])
			dayOQPClose, _ := strconv.Atoi(row[11])
			dayOQPChange, _ := strconv.Atoi(row[12])
			dayIVPercent, _ := strconv.Atoi(row[13])
			dayVolume, _ := strconv.Atoi(row[14])
			contractHigh, _ := strconv.Atoi(row[15])
			contractLow, _ := strconv.Atoi(row[16])
			volume, _ := strconv.Atoi(row[17])
			openInterest, _ := strconv.Atoi(row[18])
			changeInOpenInterest, _ := strconv.Atoi(row[19])
			entry := DailyMarketReport{
				ContractMonth: row[0],
				StrikePrice:   uint(strikePrice),
				Type:          row[2],
				AfterHoursOpeningPrice: uint(afterHoursOpeningPrice),
				AfterHoursDailyHigh:    uint(afterHoursDailyHigh),
				AfterHoursDailyLow:     uint(afterHoursDailyLow),
				AfterHoursClosePrice:   uint(afterHoursClosePrice),
				AfterHoursVolume:       uint(afterHoursVolume),
				DayOpeningPrice:        uint(dayOpeningPrice),
				DayDailyHigh:           uint(dayDailyHigh),
				DayDailyLow:            uint(dayDailyLow),
				DayOQPClose:            uint(dayOQPClose),
				DayOQPChange:           dayOQPChange,
				DayIVPercent:           dayIVPercent,
				DayVolume:              uint(dayVolume),
				ContractHigh:           uint(contractHigh),
				ContractLow:            uint(contractLow),
				Volume:                 uint(volume),
				OpenInterest:           uint(openInterest),
				ChangeInOpenInterest:   changeInOpenInterest,
			}
			if _, err := entry.ContractMonthTime(); err == nil {
				// Should be a proper line if contract month was parsed correctly
				dailyMarketReport = append(dailyMarketReport, entry)
			}
		}
	}

	return dailyMarketReport, nil
}
