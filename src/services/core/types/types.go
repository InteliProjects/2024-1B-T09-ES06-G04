package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) (int, error)
	DeleteUserByID(id int) error
	UpdateUser(user User) error
}

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	CompanyName  string `json:"company_name"`
	Office       string `json:"office"`
	LinkedinLink string `json:"linkedin_link"`
	Interest	string	`json:"interest"`
	Image     string `json:"image"`
}

type RegisterUserPayload struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	CompanyName  string `json:"company_name" validate:"required"`
	Office       string `json:"office" validate:"required"`
	LinkedinLink string `json:"linkedin_link" validate:"required"`
	Interest	string	`json:"interest"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
