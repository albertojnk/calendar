package models

type CalendarSchedule struct {
	Weekday int
	Start   string
	Finish  string
}

type Calendar struct {
	ID           int
	UserID       int
	CalendarName string
	Color        string
	Schedule     []CalendarSchedule
}
