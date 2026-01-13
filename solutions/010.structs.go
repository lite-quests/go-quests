package solutions

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func NewUser(id int, name, email string, age int) *User {
	// TODO: implement
	// Read README.md for the instructions
	newUser := User{id, name, email, age}
	return &newUser
}

func (u *User) IsAdult() bool {
	// TODO: implement
	// Read README.md for the instructions
	if u.Age >= 18 {
		return true
	}
	return false
}

func (u User) DisplayName() string {
	// TODO: implement
	// Read README.md for the instructions

	return u.Name + " <" + u.Email + ">"
}

func (u *User) UpdateEmail(email string) {
	// TODO: implement
	// Read README.md for the instructions
	u.Email = email
}

func (u *User) Birthday() {
	// TODO: implement
	// Read README.md for the instructions
	u.Age += 1
}

func CloneUser(u User) User {
	user := User{u.ID, u.Name, u.Email, u.Age}
	return user
}
