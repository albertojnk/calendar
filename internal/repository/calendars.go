package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)

func createCalendar(ctx context.Context, userID int, calendarName, color string, schedule []models.CalendarSchedule) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Calendars (user_id, calendar_name, color, schedule) VALUES ($1, $2, $3, $4)",
		userID, calendarName, color, schedule)
	if err != nil {
		return err
	}
	fmt.Println("Calendar created successfully")
	return nil
}

func getCalendarByID(ctx context.Context, calendarID int) (models.Calendar, error) {
	var calendar models.Calendar
	var scheduleArray []models.CalendarSchedule

	err := postgres.GetDBX().QueryRow(ctx, "SELECT calendar_id, user_id, calendar_name, color, schedule FROM Calendars WHERE calendar_id = $1", calendarID).Scan(
		&calendar.ID, &calendar.UserID, &calendar.CalendarName, &calendar.Color, &scheduleArray,
	)
	if err != nil {
		return calendar, err
	}

	calendar.Schedule = scheduleArray
	return calendar, nil
}

func updateCalendar(ctx context.Context, calendarID int, calendarName, color string, schedule []models.CalendarSchedule) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE Calendars SET calendar_name = $1, color = $2, schedule = $3 WHERE calendar_id = $4",
		calendarName, color, schedule, calendarID,
	)
	if err != nil {
		return err
	}
	fmt.Println("Calendar updated successfully")
	return nil
}

func deleteCalendar(ctx context.Context, calendarID int) error {
	_, err := postgres.GetDBX().Exec(ctx, "DELETE FROM Calendars WHERE calendar_id = $1", calendarID)
	if err != nil {
		return err
	}
	fmt.Println("Calendar deleted successfully")
	return nil
}
