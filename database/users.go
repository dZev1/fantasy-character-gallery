package database

import "fmt"

// Registers a user into the database
func RegisterUser(username, email, passwordHash, csrf_token, session_token string) error {
	query := `
		INSERT INTO users
		(username, email, password_hash, csrf_token, session_token)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := db.Exec(query, username, email, passwordHash, csrf_token, session_token)
	if err != nil {
		return fmt.Errorf("could not register user: %v", err)
	}

	return nil
}

// Deletes a user from the database
func DeleteUser(username, email string) error {
	query := `
		DELETE FROM users
		WHERE username=$1 AND email=$2
	`

	_, err := db.Exec(query, username, email)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}

	return nil
}

// Changes a users password in the database
func ChangePassword(username, email, newPasswordHash string) error {
	query := `
		UPDATE users
		SET password_hash=$1
		WHERE username=$2 AND email=$3
	`

	_, err := db.Exec(query, newPasswordHash, username, email)
	if err != nil {
		return fmt.Errorf("could not change password: %v", err)
	}

	return nil
}
