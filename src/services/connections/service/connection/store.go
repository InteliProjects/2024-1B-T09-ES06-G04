package connection

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/connections/types"
)

// Connection represents a connection handler with a database instance.
type Connection struct {
	db *sql.DB
}

// NewConnection creates a new instance of Connection with the provided database instance.
func NewConnection(db *sql.DB) *Connection {
	return &Connection{db: db}
}

// CreateConnection inserts a new connection into the database.
func (s *Connection) CreateConnection(connection types.Connection) error {
	_, err := s.db.Exec("INSERT INTO connections (feedback, status, project_id, user_id) VALUES ($1, $2, $3, $4)",
		connection.Feedback, connection.Status, connection.ProjectID, connection.UserID)
	if err != nil {
		return err
	}

	return nil
}

// GetConnectionByID retrieves a connection by its ID.
func (s *Connection) GetConnectionByUserID(id int) ([]types.Connection, error) {
	query := `SELECT id, feedback, status, project_id, user_id FROM connections WHERE user_id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []types.Connection
	for rows.Next() {
		var connection types.Connection
		if err := rows.Scan(&connection.ID, &connection.Feedback, &connection.Status, &connection.ProjectID, &connection.UserID); err != nil {
			return nil, err
		}
		connections = append(connections, connection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

// UpdateConnection updates an existing connection in the database.
func (s *Connection) UpdateConnection(connection types.Connection) error {
	_, err := s.db.Exec("UPDATE connections SET feedback = $1, status = $2, project_id = $3, user_id = $4 WHERE id = $5",
		connection.Feedback, connection.Status, connection.ProjectID, connection.UserID, connection.ID)
	if err != nil {
		return err
	}

	// Check if the status is 'true' (indicating the connection is accepted)
	if connection.Status {
		// Get the owner ID of the project
		projectOwnerID, err := s.GetProjectOwnerID(connection.ProjectID)
		if err != nil {
			return err
		}

		// Create a user connection between the project owner and the connection
		err = s.CreateUserConnection(projectOwnerID, connection.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create Rating
func (s *Connection) CreateRating(rating types.Ratings) error {
	_, err := s.db.Exec("INSERT INTO ratings (rating, user_id, project_id) VALUES ($1, $2, $3)",
		rating.Rating, rating.UserID, rating.ProjectID)
	if err != nil {
		return err
	}

	return nil
}

// GetRatingByUserID recupera as avaliações (ratings) pelo ID do usuário
func (s *Connection) GetRatingByUserID(id int) ([]types.Ratings, error) {
	query := `SELECT id, rating, user_id, project_id FROM ratings WHERE user_id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []types.Ratings
	for rows.Next() {
		var rating types.Ratings
		if err := rows.Scan(&rating.ID, &rating.Rating, &rating.UserID, &rating.ProjectID); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

// GetConnectionsByID recupera as conexões pelo ID do usuário
func (s *Connection) GetConnectionsByID(id int) ([]types.Connection, error) {
	query := "SELECT id, feedback, status, project_id, user_id FROM connections WHERE id = $1"

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []types.Connection
	for rows.Next() {
		var connection types.Connection
		if err := rows.Scan(&connection.ID, &connection.Feedback, &connection.Status, &connection.ProjectID, &connection.UserID); err != nil {
			return nil, err
		}
		connections = append(connections, connection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

// GetAcceptedConnectionByUserID retrieves accepted connections by user ID, including full project details and user information.
func (s *Connection) GetAcceptedConnectionByUserID(id int) ([]types.Connection, error) {
	query := `
	SELECT c.id, c.feedback, c.status, c.project_id, c.user_id,
		p.name AS project_name, p.description, p.macro_setor, p.micro_setor, p.image_link, p.user_id AS project_user_id,
		u.name AS user_name, u.company_name, u.email, u.office, u.linkedin_link, u.interest, u.image AS user_image,
		pu.name AS project_user_name
	FROM connections c
	LEFT JOIN user_connections uc ON c.id = uc.connection_id
	LEFT JOIN projects p ON c.project_id = p.id
	LEFT JOIN users u ON c.user_id = u.id
	LEFT JOIN users pu ON p.user_id = pu.id
	WHERE (c.user_id = $1 OR uc.user_id = $1) AND c.status = true;
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []types.Connection
	for rows.Next() {
		var connection types.Connection
		if err := rows.Scan(&connection.ID, &connection.Feedback, &connection.Status, &connection.ProjectID,
			&connection.UserID, &connection.ProjectName, &connection.Description,
			&connection.MacroSetor, &connection.MicroSetor, &connection.ImageLink,
			&connection.ProjectUserID, &connection.UserName, &connection.CompanyName,
			&connection.Email, &connection.Office, &connection.LinkedinLink,
			&connection.Interest, &connection.UserImage, &connection.ProjectUserName); err != nil {
			return nil, err
		}
		connections = append(connections, connection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

// CreateUserConnection cria um registro na tabela user_connections
func (s *Connection) CreateUserConnection(userID, connectionID int) error {
	_, err := s.db.Exec("INSERT INTO user_connections (user_id, connection_id) VALUES ($1, $2)", userID, connectionID)
	if err != nil {
		return err
	}
	return nil
}

// GetProjectOwnerID retorna o ID do dono do projeto com base no ID do projeto
func (s *Connection) GetProjectOwnerID(projectID int) (int, error) {
	var userID int
	query := "SELECT user_id FROM projects WHERE id = $1"
	err := s.db.QueryRow(query, projectID).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
