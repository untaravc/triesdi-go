package user_repository

type User struct {
	Id       int
	Email    string
	Password string
}

func (User) TableName() string {
	return "users"
}
