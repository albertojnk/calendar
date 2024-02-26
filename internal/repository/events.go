package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)

func createEvent(ctx context.Context, userID int, title, description, start, end, location, visibility string) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Events (user_id, title, description, start_datetime, end_datetime, location, visibility) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userID, title, description, start, end, location, visibility)
	if err != nil {
		return err
	}
	fmt.Println("Event created successfully")
	return nil
}

func getEventByID(ctx context.Context, eventID int) (models.Event, error) {
	var event models.Event
	err := postgres.GetDBX().QueryRow(ctx, "SELECT event_id, user_id, title, description, start_datetime, end_datetime, location, visibility FROM Events WHERE event_id = $1", eventID).Scan(
		&event.ID, &event.UserID, &event.Title, &event.Description, &event.Start, &event.End, &event.Location, &event.Visibility,
	)
	if err != nil {
		return event, err
	}
	return event, nil
}

func updateEvent(ctx context.Context, eventID int, title, description, start, end, location, visibility string) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE Events SET title = $1, description = $2, start_datetime = $3, end_datetime = $4, location = $5, visibility = $6 WHERE event_id = $7",
		title, description, start, end, location, visibility, eventID,
	)
	if err != nil {
		return err
	}
	fmt.Println("Event updated successfully")
	return nil
}

func deleteEvent(ctx context.Context, eventID int) error {
	_, err := postgres.GetDBX().Exec(ctx, "DELETE FROM Events WHERE event_id = $1", eventID)
	if err != nil {
		return err
	}
	fmt.Println("Event deleted successfully")
	return nil
}
