package timegen

import (
	"time"
)

// TimeGenerator 現在日時取得
type TimeGenerator interface {
	Now() time.Time
	Location(date time.Time) time.Time
	ParseInLocation(dateText string) time.Time
}

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
func (*Time) Location(date time.Time) time.Time {
	loc := getLoadLocation()

	return date.In(loc)
}

// Now 現在時刻を取得する
func (*Time) Now() time.Time {
	loc := getLoadLocation()
	return time.Now().In(loc)
}
