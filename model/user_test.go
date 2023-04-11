package model

import (
	"reflect"
	"testing"
)

func TestUsers_GetByID(t *testing.T) {
	users := Users{
		{ID: 1, Name: Name{Firstname: "User1"}},
		{ID: 2, Name: Name{Firstname: "User2"}},
		{ID: 3, Name: Name{Firstname: "User3"}},
		{ID: 4, Name: Name{Firstname: "User4"}},
	}
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		users Users
		args  args
		want  User
	}{
		{
			name:  "User exists",
			users: users,
			args:  args{id: 3},
			want:  User{ID: 3, Name: Name{Firstname: "User3"}},
		},
		{
			name:  "User doesn't exist",
			users: users,
			args:  args{id: 8},
			want:  User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.users.GetByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertCoords(t *testing.T) {
	type args struct {
		coordsStr []string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "Positive case",
			args:    args{coordsStr: []string{"-122.4194", "37.7749"}},
			want:    []float64{-2.136621598315946, 0.659296379611606},
			wantErr: false,
		},
		{
			name:    "Error case",
			args:    args{coordsStr: []string{"-122.4194", "asd"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertCoords(tt.args.coordsStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertCoords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsers_FindFurthestAwayLiving(t *testing.T) {
	users := Users{
		User{
			ID:      1,
			Name:    Name{Firstname: "User1"},
			Address: Address{Geolocation: Geolocation{Lat: "0.0", Long: "0.0"}},
		},
		User{
			ID:      2,
			Name:    Name{Firstname: "User2"},
			Address: Address{Geolocation: Geolocation{Lat: "5.0", Long: "5.0"}},
		},
		User{
			ID:      3,
			Name:    Name{Firstname: "User3"},
			Address: Address{Geolocation: Geolocation{Lat: "100.0", Long: "100.0"}},
		},
	}
	wrongCoordsUser := User{
		ID:      4,
		Name:    Name{Firstname: "WrongUser"},
		Address: Address{Geolocation: Geolocation{Lat: "asdf", Long: "asdf"}},
	}
	tests := []struct {
		name  string
		users Users
		want  Users
	}{
		{
			name:  "Empty users",
			users: Users{},
			want:  Users{},
		},
		{
			name:  "Positive case",
			users: users,
			want:  Users{users[0], users[2]},
		},
		{
			name:  "Wrong coords value",
			users: append(users, wrongCoordsUser),
			want:  Users{users[0], users[2]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.users.FindFurthestAwayLiving(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.FindFurthestAwayLiving() = %v, want %v", got, tt.want)
			}
		})
	}
}
