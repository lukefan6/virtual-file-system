package services

import (
	"reflect"
	"testing"
	"time"
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
				CreatedBy:   "Luke",
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

			// Workaround for ignoring comparison for time.Time
			if tt.want != nil {
				tt.want.CreatedAt = time.Time{}
			}

			if got != nil {
				got.CreatedAt = time.Time{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FolderServiceImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolderServiceImpl_Delete(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
	}
	type args struct {
		id        int
		deletedBy string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "01. it should delete existing folder without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				id:        1001,
				deletedBy: "Luke",
			},
			wantErr: false,
		},
		{
			name: "02. it should return error if folder does not exist.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
				},
			},
			args: args{
				id:        9999,
				deletedBy: "Luke",
			},
			wantErr: true,
		},
		{
			name: "03. it should return error if folder owner does not match.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
			},
			args: args{
				id:        1001,
				deletedBy: "Mark",
			},
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
			if err := service.Delete(tt.args.id, tt.args.deletedBy); (err != nil) != tt.wantErr {
				t.Errorf("FolderServiceImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolderServiceImpl_GetAll(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
		initKey int
	}
	type args struct {
		username  string
		sortBy    string
		sortOrder string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Folder
		wantErr bool
	}{
		// TODO: Username is taken for the purpose of authentication.
		// TODO: “Warning - empty folders”
		{
			name: "01. it should return all folders in the system without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
			},
			args: args{
				username: "Luke",
			},
			want: []models.Folder{
				{Name: "Testing", CreatedBy: "Mark"},
				{Name: "Work", CreatedBy: "Luke"},
			},
			wantErr: false,
		},
		{
			name: "02. it should return all folders order by name ascending without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
			},
			args: args{
				username:  "Luke",
				sortBy:    "sort_name",
				sortOrder: "asc",
			},
			want: []models.Folder{
				{Name: "Testing", CreatedBy: "Mark"},
				{Name: "Work", CreatedBy: "Luke"},
			},
			wantErr: false,
		},
		{
			name: "03. it should return all folders order by name descending without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
			},
			args: args{
				username:  "Luke",
				sortBy:    "sort_name",
				sortOrder: "dsc",
			},
			want: []models.Folder{
				{Name: "Work", CreatedBy: "Luke"},
				{Name: "Testing", CreatedBy: "Mark"},
			},
			wantErr: false,
		},
		{
			name: "04. it should return all folders order by created time in ascending order without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke", CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					1002: {Name: "Testing", CreatedBy: "Mark", CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
					1003: {Name: "Boss", CreatedBy: "April", CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				},
				users: map[string]models.User{
					"luke":  {Name: "Luke"},
					"mark":  {Name: "Mark"},
					"april": {Name: "April"},
				},
			},
			args: args{
				username:  "Luke",
				sortBy:    "sort_time",
				sortOrder: "asc",
			},
			want: []models.Folder{
				{Name: "Work", CreatedBy: "Luke"},
				{Name: "Boss", CreatedBy: "April"},
				{Name: "Testing", CreatedBy: "Mark"},
			},
			wantErr: false,
		},
		{
			name: "05. it should return all folders order by created time in descending order without errors.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke", CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					1002: {Name: "Testing", CreatedBy: "Mark", CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
					1003: {Name: "Boss", CreatedBy: "April", CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				},
				users: map[string]models.User{
					"luke":  {Name: "Luke"},
					"mark":  {Name: "Mark"},
					"april": {Name: "April"},
				},
			},
			args: args{
				username:  "Luke",
				sortBy:    "sort_time",
				sortOrder: "dsc",
			},
			want: []models.Folder{
				{Name: "Testing", CreatedBy: "Mark"},
				{Name: "Boss", CreatedBy: "April"},
				{Name: "Work", CreatedBy: "Luke"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &FolderServiceImpl{
				folders: tt.fields.folders,
				userService: &UserServiceImpl{
					users: tt.fields.users,
				},
				initKey: tt.fields.initKey,
			}
			got, err := service.GetAll(tt.args.username, tt.args.sortBy, tt.args.sortOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("FolderServiceImpl.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Workaround for ignoring comparison for time.Time
			if tt.want != nil {
				for i := 0; i < len(tt.want); i++ {
					tt.want[i].CreatedAt = time.Time{}
				}
			}

			if got != nil {
				for i := 0; i < len(got); i++ {
					got[i].CreatedAt = time.Time{}
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FolderServiceImpl.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolderServiceImpl_Rename(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
	}
	type args struct {
		id        int
		name      string
		renamedBy string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "01. it should rename folder without error.",
			fields: fields{
				folders: map[int]models.Folder{1001: {Name: "Work"}},
				users:   map[string]models.User{"luke": {Name: "Luke"}},
			},
			args: args{
				id:        1001,
				name:      "Work2",
				renamedBy: "Luke",
			},
			wantErr: false,
		},
		{
			name: "02. it should return error if folder not found.",
			fields: fields{
				folders: map[int]models.Folder{1001: {Name: "Work"}},
				users:   map[string]models.User{"luke": {Name: "Luke"}},
			},
			args: args{
				id:        1002,
				name:      "Work2",
				renamedBy: "Luke",
			},
			wantErr: true,
		},
		{
			name: "03. it should return error if user not found.",
			fields: fields{
				folders: map[int]models.Folder{1001: {Name: "Work"}},
				users:   map[string]models.User{"luke": {Name: "Luke"}},
			},
			args: args{
				id:        1001,
				name:      "Work2",
				renamedBy: "Mark",
			},
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
			if err := service.Rename(tt.args.id, tt.args.name, tt.args.renamedBy); (err != nil) != tt.wantErr {
				t.Errorf("FolderServiceImpl.Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
