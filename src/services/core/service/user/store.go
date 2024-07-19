package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Inteli-College/2024-1B-T09-ES06-G04/core/types"
)

// Store struct represents the user data store with a database connection
type Store struct {
	db *sql.DB
}

// NewStore creates a new instance of Store with the provided database connection
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetUserByEmail retrieves a user from the database by their email address
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// scanRowsIntoUser scans the rows returned by a query into a User struct
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.CompanyName,
		&user.Email,
		&user.Password,
		&user.Office,
		&user.LinkedinLink,
		&user.Interest,
		&user.Image,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user from the database by their ID.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT* FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// CreateUser creates a new user in the database and returns their ID
func (s *Store) CreateUser(user types.User) (int, error) {
	var userID int
	err := s.db.QueryRow("INSERT INTO users (name, company_name, email, password, office, linkedin_link, interest, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", user.Name, user.CompanyName, user.Email, user.Password, user.Office, user.LinkedinLink, user.Interest, user.Image).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// DeleteUserByID deletes a user from the database by their ID
func (s *Store) DeleteUserByID(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	
	return nil
}

// UpdateUser updates a user's information in the database
func (s *Store) UpdateUser(user types.User) error {
	query := "UPDATE users SET "
	parts := []string{}
	args := []interface{}{}
	argId := 1

	if user.Name != "" {
			parts = append(parts, fmt.Sprintf("name = $%d", argId))
			args = append(args, user.Name)
			argId++
	}
	if user.CompanyName != "" {
			parts = append(parts, fmt.Sprintf("company_name = $%d", argId))
			args = append(args, user.CompanyName)
			argId++
	}
	if user.Email != "" {
			parts = append(parts, fmt.Sprintf("email = $%d", argId))
			args = append(args, user.Email)
			argId++
	}

	if user.Interest != "" {
			parts = append(parts, fmt.Sprintf("interest = $%d", argId))
			args = append(args, user.Interest)
			argId++
	}

	if user.Image != "" {
			parts = append(parts, fmt.Sprintf("image = $%d", argId))
			args = append(args, user.Image)
			argId++
	}

	if len(parts) == 0 {
			return nil
	}

	query += strings.Join(parts, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", argId)
	args = append(args, user.ID)

	_, err := s.db.Exec(query, args...)
	if err != nil {
			return err
	}

	return nil
}
