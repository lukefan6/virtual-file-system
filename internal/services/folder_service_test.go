package services

import (
	"reflect"
	"testing"
	"virtual-file-system/internal/models"
)

func TestFolderServiceImpl_Create(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
	}
	type args struct {
		name      string
		createdBy string
		desc      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Folder
		wantErr bool
	}{
		// TODO: Add edge test cases like empty string, all whitespaces, etc.
		// TODO: Verify internal storage after test execution.
		{
			name: "01. it should add new folder when given user exists while folder does not.",
			fields: fields{
				folders: map[int]models.Folder{},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				name:      "Work",
				createdBy: "Luke",
				desc:      "The working files and necessary files are here",
			},
			want: &models.Folder{
				ID:          1001,
				Name:        "Work",
				Description: "The working files and necessary files are here",
				Files:       []models.File{},
			},
			wantErr: false,
		},
		{
			name: "02. it should return error when given user does not exist.",
			fields: fields{
				folders: map[int]models.Folder{},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				name:      "Work",
				createdBy: "abc",
				desc:      "The working files and necessary files are here",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "03. it should return error when given folder name already exist.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				name:      "Work",
				createdBy: "Luke",
				desc:      "The working files and necessary files are here",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "04. it should return error when given folder name already exist using case-insensitive comparison.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				name:      "work",
				createdBy: "Luke",
				desc:      "The working files and necessary files are here",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &FolderServiceImpl{
				folders: tt.fields.folders,
				userService: &UserServiceImpl{
					users: tt.fields.users,
				},
				initKey: 1001,
			}
			got, err := service.Create(tt.args.name, tt.args.createdBy, tt.args.desc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FolderServiceImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FolderServiceImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
