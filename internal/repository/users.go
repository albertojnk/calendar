package repository

import (
	"context"
	"fmt"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/internal/models"
)


func createUser(ctx context.Context,username, email string, documentNumber, phone int64, passwordHash string) error {
	_, err := postgres.GetDBX().Exec(ctx, "INSERT INTO Users (username, email, document_number, phone, password_hash) VALUES ($1, $2, $3, $4, $5)",
		username, email, documentNumber, phone, passwordHash)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully")
	return nil
}

func getUserByID(ctx context.Context, userID int) (models.User, error) {
	var user models.User
	err := postgres.GetDBX().QueryRow(ctx, "SELECT user_id, username, email, document_number, phone, password_hash FROM Users WHERE user_id = $1", userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.DocumentNumber, &user.Phone, &user.PasswordHash,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func updateUser(ctx context.Context, userID int, username, email string, documentNumber, phone int64, passwordHash string) error {
	_, err := postgres.GetDBX().Exec(ctx, "UPDATE Users SET username = $1, email = $2, document_number = $3, phone = $4, password_hash = $5 WHERE user_id = $6",
		username, email, documentNumber, phone, passwordHash, userID,
	)
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
