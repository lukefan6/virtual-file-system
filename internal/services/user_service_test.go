package services

import (
	"testing"
	"virtual-file-system/internal/models"
)

func TestUserServiceImpl_Register(t *testing.T) {
	type fields struct {
		users map[string]models.User
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add edge test cases like empty string, all whitespaces, etc.
		// TODO: Verify internal user storage after test execution.
		{
			name:    "01. it should add new user when the users are empty",
			fields:  fields{map[string]models.User{}},
			args:    args{"Luke"},
			wantErr: false,
		},
		{
			name: "02. it should return error if a duplicated user name is given",
			fields: fields{
				map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args:    args{"Luke"},
			wantErr: true,
		},
		{
			name: "03. it should return error if the given user name exists using case-insensitive comparison",
			fields: fields{
				map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args:    args{"luke"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserServiceImpl{
				users: tt.fields.users,
			}
			if err := service.Register(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UserServiceImpl.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
