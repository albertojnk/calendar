package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)

func createEventParticipant(ctx context.Context, eventID, userID int, status string) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Event_Participants (event_id, user_id, status) VALUES ($1, $2, $3)",
		eventID, userID, status)
	if err != nil {
		return err
	}
	fmt.Println("Event Participant created successfully")
	return nil
}

func getEventParticipant(ctx context.Context, eventID, userID int) (models.EventParticipant, error) {
	var participant models.EventParticipant
	err := postgres.GetDBX().QueryRow(ctx, "SELECT event_id, user_id, status FROM Event_Participants WHERE event_id = $1 AND user_id = $2", eventID, userID).Scan(
		&participant.EventID, &participant.UserID, &participant.Status,
	)
	if err != nil {
		return participant, err
	}
	return participant, nil
}

func updateEventParticipant(ctx context.Context, eventID, userID int, status string) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE Event_Participants SET status = $1 WHERE event_id = $2 AND user_id = $3",
		status, eventID, userID)
	if err != nil {
		return err
	}
	fmt.Println("Event Participant updated successfully")
	return nil
}

func deleteEventParticipant(ctx context.Context, eventID, userID int) error {
	_, err := postgres.GetDBX().Exec(ctx, "DELETE FROM Event_Participants WHERE event_id = $1 AND user_id = $2", eventID, userID)
	if err != nil {
		return err
	}
	fmt.Println("Event Participant deleted successfully")
	return nil
}
