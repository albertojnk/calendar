package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/albertojnk/calendar/adapters/postgres"
	"github.com/albertojnk/calendar/environment"
)

func TestCreateUser(t *testing.T) {

	environment.InitEnv("../../.env")
	postgres.InitDB()

	type args struct {
		ctx            context.Context
		username       string
		email          string
		documentNumber int64
		phone          int64
		passwordHash   string
	}
	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:            context.Background(),
				username:       "alberto",
				email:          "albertojanicke@live.com",
				documentNumber: 45500449850,
				phone:          11995708997,
				passwordHash:   "abc@123",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := createUser(tt.args.ctx, tt.args.username, tt.args.email, tt.args.documentNumber, tt.args.phone, tt.args.passwordHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("BondEntity_t.GetOverlapingBonds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("BondEntity_t.GetOverlapingBonds() = %v, want %v", err, tt.want)
			}
		})
	}
}
