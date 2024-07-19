package types

type ProjectStore interface {
	CreateProject(Project) error
	GetProjects() ([]Project, error)
	GetProjectsByUserID(id int) ([]Project, error)
	GetProjectByID(id int) (*Project, error)
	UpdateProject(Project) error
	DeleteProject(id int) error
}

type UserDetails struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MacroSetor  string `json:"macro_setor"`
	MicroSetor  string `json:"micro_setor"`
	ImageLink   string `json:"image_link"`
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`  
  	UserCompany string `json:"user_company"` 
}

type RegisterProjectPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	MacroSetor  string `json:"macro_setor" validate:"required"`
	MicroSetor  string `json:"micro_setor" validate:"required"`
	ImageLink   string `json:"image_link"`
	UserId      int    `json:"user_id" validate:"required"`
}

type UpdateProjectPayload struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	MacroSetor  string `json:"macro_setor" validate:"required"`
	MicroSetor  string `json:"micro_setor" validate:"required"`
	ImageLink   string `json:"image_link"`
	UserId      int    `json:"user_id" validate:"required"`
}
