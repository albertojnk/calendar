package models

type CalendarSchedule struct {
	ID         int
	CalendarID int
	Weekday    int
	Start      int
	Finish     int
}

type Calendar struct {
	ID           int
	UserID       int
	CalendarName string
	Color        string
	Schedule     []CalendarSchedule
}
