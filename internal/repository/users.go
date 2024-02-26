package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)

func createUser(ctx context.Context, username, email, passwordHash string) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Users (username, email, password_hash) VALUES ($1, $2, $3)",
		username, email, passwordHash)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully")
	return nil
}

func getUserByID(ctx context.Context, userID int) (models.User, error) {
	var user models.User
	err := postgres.GetDBX().QueryRow(ctx, "SELECT user_id, username, email FROM Users WHERE user_id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func updateUser(ctx context.Context, userID int, username, email string) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE Users SET username = $1, email = $2 WHERE user_id = $3", username, email, userID)
	if err != nil {
		return err
	}
	fmt.Println("User updated successfully")
	return nil
}

func deleteUser(ctx context.Context, userID int) error {
	_, err := postgres.GetDBX().Exec(ctx, "DELETE FROM Users WHERE user_id = $1", userID)
	if err != nil {
		return err
	}
	fmt.Println("User deleted successfully")
	return nil
}
