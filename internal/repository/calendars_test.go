package repository

import (
	"context"
	"testing"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/environment"
	"github.com/albertojnk/calendar/internal/models"
)

func Test_createCalendar(t *testing.T) {

	environment.InitEnv("../../.env")
	postgres.InitDB()

	type args struct {
		ctx          context.Context
		userID       int
		calendarName string
		color        string
		schedule     []models.CalendarSchedule
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				calendarName: "calendartest",
				color:        "#0981D1",
				schedule: []models.CalendarSchedule{
					{
						Weekday: 2,
						Start:   8,
						Finish:  18,
					},
					{
						Weekday: 3,
						Start:   8,
						Finish:  18,
					},
					{
						Weekday: 4,
						Start:   8,
						Finish:  18,
					},
					{
						Weekday: 5,
						Start:   10,
						Finish:  14,
					},
					{
						Weekday: 6,
						Start:   8,
						Finish:  18,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createCalendar(tt.args.ctx, tt.args.userID, tt.args.calendarName, tt.args.color, tt.args.schedule); (err != nil) != tt.wantErr {
				t.Errorf("createCalendar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
