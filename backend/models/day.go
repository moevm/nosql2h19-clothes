package models

type Weekday int

const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	return names[day]
}

func (day Weekday) Weekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

type Day struct {
	Id         int64      `json:"id"`
	DayOfWeek  Weekday    `json:"dayOfWeek"`
	StartTime  int64      `json:"startTime"`
	CloseTime  int64      `json:"closeTime"`
	Categories []Category `json:"categories"`
}

func CreateDay(c Day) int64 {
	var uid int64

	return uid
}

func UpdateDay(d Day) bool {
	return true
}

func DeleteDay(d Day) bool {
	return true
}

func GetDays() []Day {
	var ds []Day

	return ds
}

func GetDayById(id int64) *Day {
	var d Day
	return &d
}

func GetDaysByDayOfWeek(dayOfWeek string) []Day {
	var d []Day
	return d
}
