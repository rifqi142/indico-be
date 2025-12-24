package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ReadableTime struct {
	time.Time
}

func (rt ReadableTime) MarshalJSON() ([]byte, error) {
	if rt.Time.IsZero() {
		return []byte("null"), nil
	}
	
	formatted := FormatToIndonesian(rt.Time)
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}

func (rt *ReadableTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	
	str := string(data)
	if len(str) > 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}
	
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	
	var err error
	for _, format := range formats {
		rt.Time, err = time.Parse(format, str)
		if err == nil {
			return nil
		}
	}
	
	return err
}

func (rt *ReadableTime) Scan(value interface{}) error {
	if value == nil {
		rt.Time = time.Time{}
		return nil
	}
	
	if t, ok := value.(time.Time); ok {
		rt.Time = t
		return nil
	}
	
	return fmt.Errorf("cannot scan %T into ReadableTime", value)
}

func (rt ReadableTime) Value() (driver.Value, error) {
	if rt.Time.IsZero() {
		return nil, nil
	}
	return rt.Time, nil
}

func FormatToIndonesian(t time.Time) string {
	dayNames := map[time.Weekday]string{
		time.Sunday:    "Minggu",
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jumat",
		time.Saturday:  "Sabtu",
	}
	
	monthNames := map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Maret",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "Juni",
		time.July:      "Juli",
		time.August:    "Agustus",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Desember",
	}
	
	// Get WIB timezone
	wib := time.FixedZone("WIB", 7*60*60)
	localTime := t.In(wib)
	
	dayName := dayNames[localTime.Weekday()]
	day := localTime.Day()
	monthName := monthNames[localTime.Month()]
	year := localTime.Year()
	
	return fmt.Sprintf("%s, %d %s %d", dayName, day, monthName, year)
}

func FormatToIndonesianWithTime(t time.Time) string {
	wib := time.FixedZone("WIB", 7*60*60)
	localTime := t.In(wib)
	
	dateStr := FormatToIndonesian(t)
	timeStr := localTime.Format("15:04:05")
	
	return fmt.Sprintf("%s pukul %s WIB", dateStr, timeStr)
}

func NewReadableTime(t time.Time) ReadableTime {
	return ReadableTime{Time: t}
}
