package data

type Repository interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(u User) error
	DeleteByID(id int) error
	Insert(user User) (int, error)
	ResetPassword(password string, u User) error
	PasswordMatches(plainText string, u User) (bool, error)
}
