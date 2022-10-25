package domain

type UserProfile struct {
	UserId    UserId     `json:"userId"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	Gender    Gender     `json:"gender"`
	Interests []Interest `json:"interests"`
	City      string     `json:"city"`
}

func NewUserProfile(firstName string, lastName string, age int, gender Gender, interests []Interest, city string) *UserProfile {
	return &UserProfile{FirstName: firstName, LastName: lastName, Age: age, Gender: gender, Interests: interests, City: city}
}

type Interest string

type Gender int

type Login string

type UserId int64

type RegisterUser struct {
	Login     Login      `json:"login"`
	Password  Password   `json:"password"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Age       int        `json:"age"`
	Gender    Gender     `json:"gender"`
	Interests []Interest `json:"interests"`
	City      string     `json:"city"`
}
