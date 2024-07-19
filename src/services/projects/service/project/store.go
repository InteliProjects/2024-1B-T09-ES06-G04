package project

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/projects/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Function responsible for creating a project in the database
func (s *Store) CreateProject(project types.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name, description, macro_setor, micro_setor, image_link, user_id) VALUES ($1, $2, $3, $4, $5, $6)", project.Name, project.Description, project.MacroSetor, project.MicroSetor, project.ImageLink, project.UserId)
	if err != nil {
		return err
	}

	return nil
}

// Function responsible for returning up to 20 projects registered in the database
func (store *Store) GetProjects() ([]types.Project, error) {
	var projects []types.Project
	query := `
	SELECT p.id, p.name, p.description, p.macro_setor, p.micro_setor, p.image_link, p.user_id, u.name as user_name, u.company_name as user_company
	FROM projects p
	JOIN users u ON p.user_id = u.id
	LIMIT 20
	`
	rows, err := store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project types.Project

		if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.MacroSetor, &project.MicroSetor, &project.ImageLink, &project.UserId, &project.UserName, &project.UserCompany); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}



// Function responsible for returning all projects for a specific user
func (s *Store) GetProjectsByUserID(id int) ([]types.Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.Project
	for rows.Next() {
		var project types.Project
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.MacroSetor, &project.MicroSetor, &project.ImageLink, &project.UserId)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

// Function responsible for returning a specific project
func (s *Store) GetProjectByID(id int) (*types.Project, error) {
	row := s.db.QueryRow("SELECT * FROM projects WHERE id = $1", id)

	var project types.Project
	err := row.Scan(&project.Id, &project.Name, &project.Description, &project.MacroSetor, &project.MicroSetor, &project.ImageLink, &project.UserId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// Function responsible for updating a project in the database
func (s *Store) UpdateProject(project types.Project) error {
	_, err := s.db.Exec("UPDATE projects SET name = $1, description = $2, macro_setor = $3, micro_setor = $4, image_link = $5, user_id = $6 WHERE id = $7", project.Name, project.Description, project.MacroSetor, project.MicroSetor, project.ImageLink, project.UserId, project.Id)
	if err != nil {
		return err
	}

	return nil
}

// Function responsible for deleting a project from the database
func (s *Store) DeleteProject(id int) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
