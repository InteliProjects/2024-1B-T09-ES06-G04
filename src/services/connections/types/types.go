package types

type ConnectionStore interface {
	CreateConnection(Connection) error
	//GetAcceptedConnections(id int, idConnection int) ([]Connection, error)
	GetConnectionByUserID(id int) ([]Connection, error)
	UpdateConnection(Connection) error
	GetRatingByUserID(id int)([]Ratings, error)
	CreateRating(rating Ratings)error
	GetAcceptedConnectionByUserID(id int)([]Connection, error)
	GetConnectionsByID(id int) ([]Connection, error)
	CreateUserConnection(userID, connectionID int) error 
	GetProjectOwnerID(projectID int) (int, error)  
}

// type Connection struct {
// 	ID           int	 `json:"id"`
// 	Feedback     string  `json:"feedback"`
// 	Status 		 bool 	 `json:"status"`
// 	ProjectID	 int	 `json:"project_id"`
// 	UserID	int	`json:"user_id"`
// }

// type Connection struct {
// 	ID           int    `json:"id"`
// 	Feedback     string `json:"feedback"`
// 	Status       bool   `json:"status"`
// 	UserID       int    `json:"user_id"`
// 	ProjectID    int    `json:"project_id"`
// 	ProjectName  string `json:"project_name"`
// 	Description  string `json:"description"`
// 	MacroSetor   string `json:"macro_setor"`
// 	MicroSetor   string `json:"micro_setor"`
// 	ImageLink    string `json:"image_link"`
// 	ProjectUserID int   `json:"project_user_id"`
// }

type Connection struct {
	ID            int    `json:"id"`
	Feedback      string `json:"feedback"`
	Status        bool   `json:"status"`
	UserID        int    `json:"user_id"`
	ProjectID     int    `json:"project_id"`
	ProjectName   string `json:"project_name"`
	Description   string `json:"description"`
	MacroSetor    string `json:"macro_setor"`
	MicroSetor    string `json:"micro_setor"`
	ImageLink     string `json:"image_link"`
	ProjectUserID int    `json:"project_user_id"`
	ProjectUserName string `json:"project_user_name"`
	UserName      string `json:"user_name"`
	CompanyName   string `json:"company_name"`
	Email         string `json:"email"`
	Office        string `json:"office"`
	LinkedinLink  string `json:"linkedin_link"`
	Interest      string `json:"interest"`
	UserImage     string `json:"user_image"`
}

type UserConnection struct{
	ID				int		 `json:"id"`
	UserID			int		 `json:"user_id"`
	ConnectionID	int		 `json:"connection_id"`
}

type CreateConnectionPayload struct{
	Feedback     string  `json:"feedback"`
	Status 		 bool 	 `json:"status"`
	ProjectID	 int	 `json:"project_id"`
	UserID	int	`json:"user_id"`
}

type UpdateConnectionPayload struct{
	ID           int	 `json:"id"`
	Feedback     string  `json:"feedback"`
	Status 		 bool 	 `json:"status"`
	ProjectID	 int	 `json:"project_id"`
	UserID	int	`json:"user_id"`

}

type CreateRatingPayload struct{
	Rating	int	`json:"rating"`
	UserID	int	`json:"user_id"`
	ProjectID	 int	 `json:"project_id"`
}

type Ratings struct {
	ID	int	`json:"id"`
	Rating	int	`json:"rating"`
	UserID	int	`json:"user_id"`
	ProjectID	 int	 `json:"project_id"`
}

type ConnectionDetails struct {
	UserID         int    `json:"user_id"`
	ConnectionID   int    `json:"connection_id"`
	Feedback       string `json:"feedback"`
	Status         bool   `json:"status"`
	ProjectID      int    `json:"project_id"`
	ProjectOwnerID int    `json:"project_owner_id"`
	UserName       string `json:"user_name"`
	UserEmail      string `json:"user_email"`
	UserCompany    string `json:"user_company"`
}

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MacroSetor  string `json:"macro_setor"`
	MicroSetor  string `json:"micro_setor"`
	ImageLink   string `json:"image_link"`
	UserID      int    `json:"user_id"`
}
