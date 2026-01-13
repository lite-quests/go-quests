package structs

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the structs Quest 🎉")
	}
	os.Exit(code)
}

/*
========================
NewUser
========================
*/

func TestNewUser(t *testing.T) {
	tests := []struct {
		id    int
		name  string
		email string
		age   int
	}{
		{1, "Mani", "mani@example.com", 20},
		{2, "Test", "test@example.com", 0},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			u := NewUser(tt.id, tt.name, tt.email, tt.age)
			if u == nil {
				t.Fatalf("NewUser returned nil")
			}

			if u.ID != tt.id || u.Name != tt.name || u.Email != tt.email || u.Age != tt.age {
				t.Fatalf(
					"NewUser() = %+v, want ID=%d Name=%s Email=%s Age=%d",
					u, tt.id, tt.name, tt.email, tt.age,
				)
			}
		})
	}
}

/*
========================
IsAdult
========================
*/

func TestUserIsAdult(t *testing.T) {
	tests := []struct {
		age  int
		want bool
	}{
		{17, false},
		{18, true},
		{25, true},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			u := &User{Age: tt.age}
			if got := u.IsAdult(); got != tt.want {
				t.Fatalf("IsAdult() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
========================
DisplayName
========================
*/

func TestDisplayName(t *testing.T) {
	tests := []struct {
		u    User
		want string
	}{
		{
			User{Name: "Mani", Email: "mani@example.com"},
			"Mani <mani@example.com>",
		},
		{
			User{},
			" <>",
		},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			if got := tt.u.DisplayName(); got != tt.want {
				t.Fatalf("DisplayName() = %q, want %q", got, tt.want)
			}
		})
	}
}

/*
========================
UpdateEmail
========================
*/

func TestUpdateEmail(t *testing.T) {
	tests := []struct {
		start    string
		newEmail string
	}{
		{"old@example.com", "new@example.com"},
		{"x@example.com", ""},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			u := &User{Email: tt.start}
			u.UpdateEmail(tt.newEmail)

			if u.Email != tt.newEmail {
				t.Fatalf("UpdateEmail() = %q, want %q", u.Email, tt.newEmail)
			}
		})
	}
}

/*
========================
Birthday
========================
*/

func TestBirthday(t *testing.T) {
	tests := []struct {
		start int
		want  int
	}{
		{20, 21},
		{0, 1},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			u := &User{Age: tt.start}
			u.Birthday()

			if u.Age != tt.want {
				t.Fatalf("Birthday() age = %d, want %d", u.Age, tt.want)
			}
		})
	}
}

/*
========================
CloneUser
========================
*/

func TestCloneUser(t *testing.T) {
	tests := []struct {
		u User
	}{
		{
			User{ID: 1, Name: "Mani", Email: "mani@example.com", Age: 20},
		},
		{
			User{},
		},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			clone := CloneUser(tt.u)

			if clone != tt.u {
				t.Fatalf("CloneUser() = %+v, want %+v", clone, tt.u)
			}

			clone.Age++
			if tt.u.Age == clone.Age {
				t.Fatalf("CloneUser() did not create an independent copy")
			}
		})
	}
}
