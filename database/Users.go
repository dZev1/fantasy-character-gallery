package database

import "fmt"

// Registers a user into the database
func RegisterUser(username string, password_hash string, csrf_token string, session_token string) error {
	query := `
		INSERT INTO users
		(username, password, csrf_token, session_token)
		VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query, username, password_hash, csrf_token, session_token)
	if err != nil {
		return fmt.Errorf("could not register user: %v", err)
	}
	
	return nil
}