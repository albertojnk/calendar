package models

type Event struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Start       string
	End         string
	Location    string
	Visibility  string
}
