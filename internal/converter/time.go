package converter

import "time"

var _ ConvertSafeFunc[string, time.Time] = ConvertStringToDate

func ConvertStringToDate(src string) (time.Time, error) {
	return time.Parse(time.DateOnly, src)
}

var _ ConvertFunc[time.Time, string] = ConvertDateToString

func ConvertDateToString(src time.Time) string {
	return src.Format("2006-01-02")
}
