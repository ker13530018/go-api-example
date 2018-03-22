package user

import (
	"reflect"
	"testing"
)

func TestUsers_find(t *testing.T) {

	all := make(map[int]Model)
	Snooker := Model{
		ID:   1,
		Name: "Snooker",
		Tel:  "0875684XXX",
	}

	all[1] = Snooker

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		u       *Users
		args    args
		want    *Model
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			u: &Users{
				LastInsertID: 0,
				All:          all,
			},
			args: args{
				id: 1,
			},
			want:    &Snooker,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_findAll(t *testing.T) {
	all := make(map[int]Model)
	Snooker := Model{
		ID:   1,
		Name: "Snooker",
		Tel:  "0875684XXX",
	}

	all[1] = Snooker

	users := Users{
		All: all,
	}

	tests := []struct {
		name    string
		u       *Users
		want    *List
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			u:       &users,
			want:    &users.All,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.findAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.findAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.findAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_store(t *testing.T) {
	all := make(map[int]Model)
	Snooker := Model{
		ID:   1,
		Name: "Snooker",
		Tel:  "0875684XXX",
	}

	all[1] = Snooker

	users := Users{
		All:          all,
		LastInsertID: 10,
	}

	type args struct {
		model *Model
	}
	tests := []struct {
		name    string
		u       *Users
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			u:    &users,
			args: args{
				model: &Model{
					Name: "Snooker",
					Tel:  "0995684XXX",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.store(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Users.store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
