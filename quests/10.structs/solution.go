package structs

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func NewUser(id int, name, email string, age int) *User {
	// TODO: implement
	// Read README.md for the instructions
	return nil
}

func (u *User) IsAdult() bool {
	// TODO: implement
	// Read README.md for the instructions
	return false
}

func (u User) DisplayName() string {
	// TODO: implement
	// Read README.md for the instructions
	return ""
}

func (u *User) UpdateEmail(email string) {
	// TODO: implement
	// Read README.md for the instructions
}

func (u *User) Birthday() {
	// TODO: implement
	// Read README.md for the instructions
}

func CloneUser(u User) User {
	return User{}
}
