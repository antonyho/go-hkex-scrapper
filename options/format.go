package options

import "time"

func FormatURLDate(d time.Time) string {
	return d.Format("060102") // Example: 22nd Mar 2019 -> 190322
}
