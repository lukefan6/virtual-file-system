package services

import (
	"reflect"
	"testing"
	"time"
	"virtual-file-system/internal/models"
)

func TestFileServiceImpl_GetAll(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
		files   map[string]models.File
	}
	type args struct {
		username  string
		folderID  int
		sortBy    string
		sortOrder string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.File
		wantErr bool
	}{
		{
			name: "01. it should return all files under given folder with default ordering.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002},
				},
			},
			args: args{
				username: "Mark",
				folderID: 1002,
			},
			want: []models.File{
				{Name: "1.png", Ext: "png", FolderID: 1002},
				{Name: "1.tc", Ext: "tc", FolderID: 1002},
			},
			wantErr: false,
		},
		{
			name: "02. it should return all files under given folder sorted by name in ascending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002},
				},
			},
			args: args{
				username: "Mark",
				folderID: 1002,
			},
			want: []models.File{
				{Name: "1.png", Ext: "png", FolderID: 1002},
				{Name: "1.tc", Ext: "tc", FolderID: 1002},
				{Name: "2.png", Ext: "png", FolderID: 1002},
			},
			wantErr: false,
		},
		{
			name: "03. it should return all files under given folder sorted by name in descending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002},
				},
			},
			args: args{
				username:  "Mark",
				folderID:  1002,
				sortBy:    "sort_name",
				sortOrder: "dsc",
			},
			want: []models.File{
				{Name: "2.png", Ext: "png", FolderID: 1002},
				{Name: "1.tc", Ext: "tc", FolderID: 1002},
				{Name: "1.png", Ext: "png", FolderID: 1002},
			},
			wantErr: false,
		},
		{
			name: "04. it should return all files under given folder sorted by created time in ascending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username:  "Mark",
				folderID:  1002,
				sortBy:    "sort_time",
				sortOrder: "asc",
			},
			want: []models.File{
				{Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
				{Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
				{Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "05. it should return all files under given folder sorted by created time in descending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username:  "Mark",
				folderID:  1002,
				sortBy:    "sort_time",
				sortOrder: "dsc",
			},
			want: []models.File{
				{Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				{Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				{Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
				{Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "06. it should return all files under given folder sorted by extension in ascending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username:  "Mark",
				folderID:  1002,
				sortBy:    "sort_extension",
				sortOrder: "asc",
			},
			want: []models.File{
				{Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				{Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
				{Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
				{Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "07. it should return all files under given folder sorted by extension in descending order.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username:  "Mark",
				folderID:  1002,
				sortBy:    "sort_extension",
				sortOrder: "dsc",
			},
			want: []models.File{
				{Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
				{Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				{Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
				{Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
			},
			wantErr: false,
		},
		{
			name: "08. it should return error if user does not exist.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username: "abc",
				folderID: 1002,
			},
			wantErr: true,
		},
		{
			name: "09. it should return empty files without error.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username: "Luke",
				folderID: 1001,
			},
			want:    []models.File{},
			wantErr: false,
		},
		{
			name: "10. it should return error when folder does not exist.",
			fields: fields{
				folders: map[int]models.Folder{
					1001: {Name: "Work", CreatedBy: "Luke"},
					1002: {Name: "Testing", CreatedBy: "Mark"},
				},
				users: map[string]models.User{
					"luke": {Name: "Luke"},
					"mark": {Name: "Mark"},
				},
				files: map[string]models.File{
					"1.tc":  {Name: "1.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 24, 0, 0, 0, 0, time.UTC)},
					"2.png": {Name: "2.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 23, 0, 0, 0, 0, time.UTC)},
					"1.png": {Name: "1.png", Ext: "png", FolderID: 1002, CreatedAt: time.Date(2021, 2, 25, 0, 0, 0, 0, time.UTC)},
					"2.tc":  {Name: "2.tc", Ext: "tc", FolderID: 1002, CreatedAt: time.Date(2021, 2, 26, 0, 0, 0, 0, time.UTC)},
				},
			},
			args: args{
				username: "Luke",
				folderID: 1003,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &FileServiceImpl{
				files: tt.fields.files,
				userService: &UserServiceImpl{
					users: tt.fields.users,
				},
				folderService: &FolderServiceImpl{
					folders: tt.fields.folders,
				},
			}
			got, err := service.GetAll(tt.args.username, tt.args.folderID, tt.args.sortBy, tt.args.sortOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileServiceImpl.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileServiceImpl.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileServiceImpl_Delete(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
		files   map[string]models.File
	}
	type args struct {
		deletedBy string
		folderID  int
		filename  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "01. it should delete file without error.",
			fields: fields{
				folders: map[int]models.Folder{1001: {Name: "Work", CreatedBy: "Luke"}},
				users:   map[string]models.User{"luke": {Name: "Luke"}},
				files:   map[string]models.File{"1.tc": {Name: "1.tc", Ext: "tc", FolderID: 1001}},
			},
			args: args{
				deletedBy: "Luke",
				folderID:  1001,
				filename:  "1.tc",
			},
			wantErr: false,
		},
		{name: "02. it should return error if folder not found."},
		{name: "03. it should return error if file not found."},
		{name: "04. it should return error if user not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &FileServiceImpl{
				files:         tt.fields.files,
				userService:   &UserServiceImpl{users: tt.fields.users},
				folderService: &FolderServiceImpl{folders: tt.fields.folders},
			}
			if err := service.Delete(tt.args.deletedBy, tt.args.folderID, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("FileServiceImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileServiceImpl_Upload(t *testing.T) {
	type fields struct {
		folders map[int]models.Folder
		users   map[string]models.User
		files   map[string]models.File
	}
	type args struct {
		createdBy string
		folderID  int
		filename  string
		desc      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "01. it should upload file without errors.",
			fields: fields{
				folders: map[int]models.Folder{1001: {Name: "Work", CreatedBy: "Luke"}},
				users:   map[string]models.User{"luke": {Name: "Luke"}},
				files:   map[string]models.File{},
			},
			args: args{
				createdBy: "Luke",
				folderID:  1001,
				filename:  "1.tc",
				desc:      "first test case for a company",
			},
			wantErr: false,
		},
		{name: "02. it should return error if folder not found."},
		{name: "03. it should return error if user not found."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &FileServiceImpl{
				files:         tt.fields.files,
				userService:   &UserServiceImpl{users: tt.fields.users},
				folderService: &FolderServiceImpl{folders: tt.fields.folders},
			}
			if err := service.Upload(tt.args.createdBy, tt.args.folderID, tt.args.filename, tt.args.desc); (err != nil) != tt.wantErr {
				t.Errorf("FileServiceImpl.Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
