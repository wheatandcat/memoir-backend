package mock_timegen

import (
	"time"
)

// Time has generating method.
type Time struct {
}

const location = "Asia/Tokyo"

func getLoadLocation() *time.Location {
	loc, _ := time.LoadLocation(location)
	return loc
}

// Location timezoneを設定する
func (*Time) ParseInLocation(dateText string) time.Time {
	loc := getLoadLocation()

	d, _ := time.ParseInLocation("2006-01-02T15:04:05", dateText, loc)
	return d
}

// Location timezoneを設定する
func (*Time) ParseInLocationTimezone(dateText string) time.Time {
	loc := getLoadLocation()

	d, _ := time.ParseInLocation("2006-01-02T15:04:05+09:00", dateText, loc)
	return d
}

// Location timezoneを設定する
func (*Time) Location(date time.Time) time.Time {
	loc := getLoadLocation()

	return date.In(loc)
}

// Now 現在時刻を取得する
func (*Time) Now() time.Time {
	loc := getLoadLocation()
	return time.Date(2020, 1, 1, 00, 00, 00, 0, loc)
}
