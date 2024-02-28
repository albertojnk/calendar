package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)

func createCalendar(ctx context.Context, userID int, calendarName, color string, schedule []models.CalendarSchedule) error {

	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Calendars (user_id, calendar_name, color) VALUES ($1, $2, $3)",
		userID, calendarName, color)
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

func createCalendarSchedule(ctx context.Context, calendarID, weekday, startTime, finishTime int) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO calendar_schedule (calendar_id, weekday, start_time, finish_time) VALUES ($1, $2, $3, $4)",
		calendarID, weekday, startTime, finishTime)
	if err != nil {
		return err
	}
	fmt.Println("Calendar Schedule created successfully")
	return nil
}

func getCalendarScheduleByID(ctx context.Context, scheduleID int) (models.CalendarSchedule, error) {
	var schedule models.CalendarSchedule

	err := postgres.GetDBX().QueryRow(ctx, "SELECT schedule_id, calendar_id, weekday, start_time, finish_time FROM calendar_schedule WHERE schedule_id = $1", scheduleID).
		Scan(&schedule.ID, &schedule.CalendarID, &schedule.Weekday, &schedule.Start, &schedule.Finish)
	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func updateCalendarSchedule(ctx context.Context, scheduleID, calendarID, weekday, startTime, finishTime int) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE calendar_schedule SET calendar_id = $1, weekday = $2, start_time = $3, finish_time = $4 WHERE schedule_id = $5",
		calendarID, weekday, startTime, finishTime, scheduleID)
	if err != nil {
		return err
	}
	fmt.Println("Calendar Schedule updated successfully")
	return nil
}

func deleteCalendarSchedule(ctx context.Context, scheduleID int) error {
	_, err := postgres.GetDBX().Exec(ctx, "DELETE FROM calendar_schedule WHERE schedule_id = $1", scheduleID)
	if err != nil {
		return err
	}
	fmt.Println("Calendar Schedule deleted successfully")
	return nil
}
